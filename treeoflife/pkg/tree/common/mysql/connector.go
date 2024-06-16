package mysql

import (
	"database/sql"
	"time"

	"github.com/cenkalti/backoff"
	"github.com/pkg/errors"
)

const dbDriver = "mysql"

type Config struct {
	MaxConnections     int
	ConnectionLifetime time.Duration
	ConnectTimeout     time.Duration // 0 means default timeout (15 seconds)
}

func NewConnector() Connector {
	return &connector{}
}

type Connector interface {
	Open(dsn DSN, cfg Config) error
	Ping() error
	Close() error
	Client() Client
}

type connector struct {
	db *sql.DB
}

type dbClient interface {
	SetMaxOpenConns(maxConnections int)
	SetConnMaxLifetime(d time.Duration)
	Ping() error
	Close() error
}

func (c *connector) Open(dsn DSN, cfg Config) error {
	db, err := sql.Open(dbDriver, dsn.String())
	if err != nil {
		return errors.Wrapf(err, "failed to open database")
	}
	err = setupDB(db, cfg)
	if err != nil {
		return errors.WithStack(err)
	}
	c.db = db
	return nil
}

func (c *connector) Ping() error {
	return c.db.Ping()
}

func (c *connector) Close() error {
	if c.db != nil {
		err := c.db.Close()
		return errors.Wrap(err, "failed to disconnect")
	}
	return errors.New("DB not initialized")
}

func (c *connector) Client() Client {
	return &client{c.db}
}

func setupDB(db dbClient, cfg Config) error {
	db.SetMaxOpenConns(cfg.MaxConnections)
	db.SetConnMaxLifetime(cfg.ConnectionLifetime)

	var err = backoff.Retry(func() error {
		tryError := db.Ping()
		return tryError
	}, newExponentialBackOff(cfg.ConnectTimeout))
	if err != nil {
		dbCloseErr := db.Close()
		if dbCloseErr != nil {
			err = errors.Wrap(err, dbCloseErr.Error())
		}
		return errors.Wrapf(err, "failed to ping database")
	}
	return nil
}

func newExponentialBackOff(timeout time.Duration) *backoff.ExponentialBackOff {
	exponentialBackOff := backoff.NewExponentialBackOff()
	const maxReconnectWaitingTime = 15 * time.Second
	if timeout != 0 {
		exponentialBackOff.MaxElapsedTime = timeout
	} else {
		exponentialBackOff.MaxElapsedTime = maxReconnectWaitingTime
	}
	return exponentialBackOff
}
