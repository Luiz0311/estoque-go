package config

import (
	"os"
	"path/filepath"
	"strconv"

	"github.com/joho/godotenv"
)

type dbConfig struct {
	DBName   string
	Host     string
	Password string
	Port     int
}

func findProjectRoot() string {
	dir, _ := os.Getwd()
	for dir != "/" {
		if _, err := os.Stat(filepath.Join(dir, ".env")); err == nil {
			return dir
		}
		dir = filepath.Dir(dir)
	}
	return "."
}

func loadEnv() error {
	root := findProjectRoot()

	if err := godotenv.Load(filepath.Join(root, ".env")); err != nil {
		return err
	}

	return nil
}

func LoadDBConfig() (dbConfig, error) {
	err := loadEnv()
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
