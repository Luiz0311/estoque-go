package repository

import (
	"database/sql"
	"time"

	"github.com/Luiz0311/estoque-go/models"
)

func (pr *ProductRepository) DeleteProduct(id int) (p models.Product, err error) {
	query := `
		UPDATE products
		SET deleted_at = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL
		RETURNING id, created_at, updated_at, deleted_at, amount, price, total_value, name, type, ean_code, available
	`

	deletedAt := time.Now()
	updatedAt := time.Now()

	err = pr.connection.QueryRow(query, deletedAt, updatedAt, id).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
		&deletedAt,
		&p.Amount,
		&p.Price,
		&p.TotalValue,
		&p.Name,
		&p.Type,
		&p.EANCode,
		&p.Available,
	)
	if err == sql.ErrNoRows {
		pr.logger.Err(models.ErrProductNotFound)
		return models.Product{}, models.ErrProductNotFound
	} else if err != nil {
		pr.logger.Err(models.ErrDelete)
		return models.Product{}, models.ErrDelete
	}

	return p, nil
}
