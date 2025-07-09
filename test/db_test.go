package test

import (
	"os"
	"testing"

	"github.com/joho/godotenv"
)

func TestEnv(t *testing.T) {
	envPath := "../.env"
	godotenv.Load(envPath)

	var (
		dbName   = os.Getenv("DBNAME")
		host     = os.Getenv("HOST")
		password = os.Getenv("PASSWORD")
		port     = os.Getenv("PORT")
	)

	if dbName == "" || host == "" || password == "" || port == "" {
		t.Error("Alguma variável de ambiente está faltando")
	}

	t.Logf(
		"\nName: %v\nHost: %v\nPassword:%v\nPort: %v",
		dbName, host, password, port,
	)
}
