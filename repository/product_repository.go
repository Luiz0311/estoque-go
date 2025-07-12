package repository

import (
	"database/sql"

	"github.com/Luiz0311/estoque-go/config"
)

type ProductRepository struct {
	connection *sql.DB
	logger     *config.Logger
}

func NewProductRepository(connection *sql.DB, logger *config.Logger) ProductRepository {
	return ProductRepository{
		connection: connection,
		logger:     logger,
	}
}
