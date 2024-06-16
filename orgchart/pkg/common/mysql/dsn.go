package mysql

import (
	"fmt"
	"strings"
)

type DSN struct {
	User          string
	Password      string
	Host          string
	Database      string
	TLSSkipVerify bool
}

func (dsn *DSN) String() string {
	options := []string{
		"charset=utf8mb4",
		"collation=utf8mb4_unicode_ci",
		"parseTime=true",
	}
	if dsn.TLSSkipVerify {
		options = append(options, "tls=skip-verify")
	}
	optionsString := strings.Join(options, "&")
	return fmt.Sprintf("%s:%s@(%s)/%s?%s", dsn.User, dsn.Password, dsn.Host, dsn.Database, optionsString)
}
