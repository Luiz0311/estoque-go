package repository

import (
	"database/sql"

	"github.com/Luiz0311/estoque-go/models"
)

func (pr *ProductRepository) GetProduct(id int) (p models.Product, err error) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, amount, price, total_value, name, type, ean_code, available
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`

	var deletedAt sql.NullTime

	err = pr.connection.QueryRow(query, id).Scan(
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
		pr.logger.Err(models.ErrSearchProduct)
		return models.Product{}, models.ErrSearchProduct
	}

	if deletedAt.Valid {
		p.DeletedAt = &deletedAt.Time
	}

	return p, nil
}
