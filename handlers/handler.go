package handlers

import (
	"database/sql"

	"github.com/Luiz0311/estoque-go/config"
)

var db *sql.DB

func InitializeHandler() {
	db = config.GetPostgres()
}
