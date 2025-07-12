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
		pr.logger.Errf("produto não encontrado: %v", err)
		return models.Product{}, err
	} else if err != nil {
		pr.logger.Errf("erro ao buscar o produto: %v", err)
		return models.Product{}, err
	}

	if deletedAt.Valid {
		p.DeletedAt = &deletedAt.Time
	}

	return p, nil
}
