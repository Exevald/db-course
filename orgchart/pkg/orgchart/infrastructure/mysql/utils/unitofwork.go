package utils

import (
	"context"
	"github.com/pkg/errors"
	"sync"
)

type UnitOfWorkFactory interface {
	NewUnitOfWork(ctx context.Context) (UnitOfWork, error)
}

type UnitOfWork interface {
	Complete(err error) error
	Client() Client
	ClientContext() ClientContext
}

type UnitOfWorkCompleteNotifier func(ctx context.Context, err error)

func NewUnitOfWorkFactory(
	connectionProvider ConnectionProvider,
	unitOfWorkCompleteNotifier UnitOfWorkCompleteNotifier,
) UnitOfWorkFactory {
	return &unitOfWorkFactory{
		connectionProvider:         connectionProvider,
		unitOfWorkCompleteNotifier: unitOfWorkCompleteNotifier,
		transactionPool:            map[context.Context]*transactionPoolEntry{},
	}
}

type unitOfWorkFactory struct {
	connectionProvider         ConnectionProvider
	mu                         sync.Mutex
	transactionPool            map[context.Context]*transactionPoolEntry
	unitOfWorkCompleteNotifier UnitOfWorkCompleteNotifier
}

type transactionPoolEntry struct {
	transaction *sharedTransaction
	count       uint
}

func (factory *unitOfWorkFactory) NewUnitOfWork(ctx context.Context) (unitOfWorkObject UnitOfWork, err error) {
	factory.withLock(func() {
		entry, ok := factory.transactionPool[ctx]
		if !ok {
			return
		}

		unitOfWorkObject = &unitOfWork{
			tx:               entry.transaction,
			completeNotifier: factory.unitOfWorkCompleteNotifier,
		}
		entry.count++
	})

	if unitOfWorkObject == nil {
		var conn TransactionalConnection
		conn, err = factory.connectionProvider.Connection(ctx)
		if err != nil {
			err = errors.WithStack(err)
			return unitOfWorkObject, err
		}

		var tx Transaction
		tx, err = conn.BeginTransaction(ctx, nil)
		if err != nil {
			if closeErr := conn.Close(); closeErr != nil {
				err = errors.Wrap(err, closeErr.Error())
			}
			return unitOfWorkObject, err
		}

		sharedTx := &sharedTransaction{
			Transaction:      tx,
			ctx:              ctx,
			conn:             conn,
			commitCallback:   factory.releaseWithCommit,
			rollbackCallback: factory.releaseWithRollback,
		}

		tx = sharedTx

		unitOfWorkObject = &unitOfWork{
			ctx:              ctx,
			tx:               tx,
			completeNotifier: factory.unitOfWorkCompleteNotifier,
		}

		factory.withLock(func() {
			factory.transactionPool[ctx] = &transactionPoolEntry{
				transaction: sharedTx,
				count:       1,
			}
		})
	}

	return unitOfWorkObject, err
}

func (factory *unitOfWorkFactory) releaseWithCommit(ctx context.Context) error {
	return factory.releaseWithCallback(ctx, func(tx *sharedTransaction) error {
		return tx.commit()
	})
}

func (factory *unitOfWorkFactory) releaseWithRollback(ctx context.Context) error {
	return factory.releaseWithCallback(ctx, func(tx *sharedTransaction) error {
		return tx.rollback()
	})
}

func (factory *unitOfWorkFactory) releaseWithCallback(ctx context.Context, f func(tx *sharedTransaction) error) (err error) {
	factory.withLock(func() {
		entry, ok := factory.transactionPool[ctx]
		if !ok {
			return
		}

		if entry.count == 1 {
			err = f(entry.transaction)
			if closeErr := entry.transaction.conn.Close(); closeErr != nil {
				err = errors.Wrap(err, closeErr.Error())
			}
			delete(factory.transactionPool, ctx)
			return
		}
		entry.count--
	})
	return
}

type unitOfWork struct {
	ctx              context.Context
	tx               Transaction
	completeNotifier UnitOfWorkCompleteNotifier
}

func (u *unitOfWork) Client() Client {
	return u.tx
}

func (u *unitOfWork) ClientContext() ClientContext {
	return u.tx
}

func (u *unitOfWork) Complete(err error) (resultErr error) {
	resultErr = err

	defer func() {
		if u.completeNotifier != nil {
			u.completeNotifier(u.ctx, resultErr)
		}
	}()

	if resultErr != nil {
		rollbackErr := u.tx.Rollback()
		if rollbackErr != nil {
			resultErr = errors.Wrap(resultErr, rollbackErr.Error())
			return
		}
		return
	}

	resultErr = errors.WithStack(u.tx.Commit())

	return resultErr
}

func (factory *unitOfWorkFactory) withLock(f func()) {
	factory.mu.Lock()
	defer factory.mu.Unlock()
	f()
}

type sharedTransaction struct {
	Transaction

	ctx              context.Context
	conn             TransactionalConnection
	commitCallback   func(ctx context.Context) error
	rollbackCallback func(ctx context.Context) error
}

func (tx *sharedTransaction) Commit() error {
	return tx.commitCallback(tx.ctx)
}

func (tx *sharedTransaction) Rollback() error {
	return tx.rollbackCallback(tx.ctx)
}

func (tx *sharedTransaction) commit() error {
	return tx.Transaction.Commit()
}

func (tx *sharedTransaction) rollback() error {
	return tx.Transaction.Rollback()
}
