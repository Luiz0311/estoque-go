package repository

import (
	"time"

	"github.com/Luiz0311/estoque-go/models"
	"github.com/Luiz0311/estoque-go/utils"
)

func (pr *ProductRepository) CreateProdct(p models.Product) (models.Product, error) {
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.EANCode = utils.GenerateEAN13()
	p.Available = p.Amount > 0
	p.TotalValue = p.Price * float64(p.Amount)

	query := `
		INSERT INTO products (name, type, ean_code, amount, price, total_value, available, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8, $9)
		RETURNING id, created_at, updated_at
	`

	err := pr.connection.QueryRow(
		query,
		p.Name,
		p.Type,
		p.EANCode,
		p.Amount,
		p.Price,
		p.TotalValue,
		p.Available,
		p.CreatedAt,
		p.UpdatedAt,
	).Scan(&p.ID, &p.CreatedAt, &p.UpdatedAt)

	if err != nil {
		pr.logger.Errf("erro ao criar o produto: %v", err)
		return models.Product{}, err
	}

	return p, nil
}
