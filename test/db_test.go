package test

import (
	"os"
	"testing"

	"github.com/Luiz0311/estoque-go/db"
	"github.com/joho/godotenv"
)

func TestPingDB(t *testing.T) {
	db, err := db.GetDB()
	if err != nil {
		t.Errorf("Falha ao instanciar base de dados")
	}
	defer db.Close()

	if err = db.Ping(); err != nil {
		t.Errorf("Falha ao conectar na base de dados")
	}
}

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
