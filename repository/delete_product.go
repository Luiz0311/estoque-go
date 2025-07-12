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
		RETURNING id, name
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
		pr.logger.Errf("produto não encontrado ou já deletado: %v", err)
		return models.Product{}, nil
	} else if err != nil {
		pr.logger.Errf("erro ao deletar produto: %v", err)
		return models.Product{}, nil
	}

	return p, nil
}
