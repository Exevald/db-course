package service

import (
	"context"

	"orgchart/pkg/orgchart/app/model"
)

type StorageProvider interface {
	BranchStorage() model.BranchStorage
	EmployeeStorage() model.EmployeeStorage
}

type UnitOfWork interface {
	StorageProvider
	Complete(err error) error
}

type UnitOfWorkFactory interface {
	NewUnitOfWork(ctx context.Context, lockName string) (UnitOfWork, error)
}

type unitOfWorkService struct {
	unitOfWorkFactory UnitOfWorkFactory
}

func (s *unitOfWorkService) executeUnitOfWork(ctx context.Context, f func(provider StorageProvider) error) (err error) {
	var unitOfWork UnitOfWork
	unitOfWork, err = s.unitOfWorkFactory.NewUnitOfWork(ctx, "")
	if err != nil {
		return err
	}
	defer func() {
		err = unitOfWork.Complete(err)
	}()
	err = f(unitOfWork)
	return err
}

func (s *unitOfWorkService) executeUnitOfWorkWithLock(
	ctx context.Context,
	lockName string,
	f func(storageProvider StorageProvider) error,
) (err error) {
	var unitOfWork UnitOfWork
	unitOfWork, err = s.unitOfWorkFactory.NewUnitOfWork(ctx, lockName)
	if err != nil {
		return err
	}
	defer func() {
		err = unitOfWork.Complete(err)
	}()
	err = f(unitOfWork)
	return err
}
