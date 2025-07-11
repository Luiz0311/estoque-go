package config

import (
	"database/sql"
)

var (
	db     *sql.DB
	logger *Logger
)

func Init() error {
	var err error

	db, err = InitializePortgres()
	if err != nil {
		return err
	}

	return nil
}

func GetPostgres() *sql.DB {
	return db
}

func GetLogger(p string) *Logger {
	logger = NewLogger(p)
	return logger
}
