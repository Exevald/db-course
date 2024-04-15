package mysql

import (
	"database/sql"
)

type Client interface {
	Query(query string, args ...interface{}) (*sql.Rows, error)
	QueryRow(query string, args ...interface{}) *sql.Row
	Exec(query string, args ...interface{}) (sql.Result, error)
}

type client struct {
	*sql.DB
}
