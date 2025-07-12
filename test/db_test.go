package test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnv(t *testing.T) {
	envPath := "../.env"
	if err := godotenv.Load(envPath); err != nil {
		t.Error("faltando o arquivo \".env\"")
		return
	}

	var (
		dbName   = os.Getenv("DBNAME")
		host     = os.Getenv("HOST")
		password = os.Getenv("PASSWORD")
		port     = os.Getenv("PORT")
	)

	if dbName == "" || host == "" || password == "" || port == "" {
		t.Error("alguma variável de ambiente está faltando")
		return
	}

	t.Logf(
		"\nName: %v\nHost: %v\nPassword:%v\nPort: %v",
		dbName, host, password, port,
	)
}
