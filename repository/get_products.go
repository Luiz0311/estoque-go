package repository

import (
	"database/sql"

	"github.com/Luiz0311/estoque-go/models"
)

func (pr *ProductRepository) GetProducts() ([]models.Product, error) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, amount, price, total_value, name, type, ean_code, available 
		FROM products 
		WHERE deleted_at IS NULL
	`

	rows, err := pr.connection.Query(query)
	if err != nil {
		pr.logger.Err(models.ErrDB)
		return []models.Product{}, models.ErrDB
	}
	defer rows.Close()

	var products []models.Product
	found := false

	for rows.Next() {
		found = true
		var p models.Product
		var deletedAt sql.NullTime

		if err := rows.Scan(
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
		); err != nil {
			pr.logger.Err(models.ErrRead)
			return []models.Product{}, models.ErrRead
		}

		if deletedAt.Valid {
			p.DeletedAt = &deletedAt.Time
		}

		products = append(products, p)
	}

	if !found {
		pr.logger.Err(models.ErrProductsNotFound)
		return []models.Product{}, models.ErrProductsNotFound
	}

	return products, nil
}
