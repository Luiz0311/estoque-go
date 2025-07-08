package db

import (
	"database/sql"
	"fmt"

	"github.com/Luiz0311/estoque-go/config"
	_ "github.com/lib/pq"
)

func GetDB() (*sql.DB, error) {
	config, err := config.LoadDBConfig()
	if err != nil {
		return nil, err
	}

	connStr := fmt.Sprintf(
		"postgres://postgres:%s@%s:%d/%s?sslmode=disable",
		config.Password, config.Host, config.Port, config.DBName,
	)

	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	return db, nil
}
