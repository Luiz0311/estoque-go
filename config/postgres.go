package config

import (
	"database/sql"
	"fmt"
	"os"
	"strconv"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

type dbConfig struct {
	DBName   string
	Host     string
	Password string
	Port     int
}

func InitializePortgres() (*sql.DB, error) {
	config, err := loadDBConfig()
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

	tableQuery, err := loadTable()
	if err != nil {
		return nil, err
	}

	_, err = db.Exec(tableQuery)
	if err != nil {
		return nil, err
	}

	return db, nil
}

func loadDBConfig() (dbConfig, error) {
	err := godotenv.Load()
	if err != nil {
		return dbConfig{}, err
	}

	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		return dbConfig{}, err
	}

	return dbConfig{
		DBName:   os.Getenv("DBNAME"),
		Host:     os.Getenv("HOST"),
		Password: os.Getenv("PASSWORD"),
		Port:     port,
	}, nil
}

func loadTable() (string, error) {
	tableStr, err := os.ReadFile("config/products_table.sql")
	if err != nil {
		return "", err
	}

	return string(tableStr), nil
}
