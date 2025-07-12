package repository

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/Luiz0311/estoque-go/models"
)

func (pr *ProductRepository) UpdateProduct(id int, data map[string]any) (p models.Product, err error) {
	if len(data) == 0 {
		pr.logger.Err("nenhum campo enviado")
		return models.Product{}, err
	}

	if amountVal, ok := data["amount"]; ok {
		if amount, ok := amountVal.(float64); ok && int(amount) == 0 {
			data["available"] = false
		}
	}

	data["updated_at"] = time.Now()

	setParts := []string{}
	values := []any{}
	i := 1

	_, hasPrice := data["price"]
	_, hasAmount := data["amount"]

	for k, v := range data {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", k, i))
		values = append(values, v)
		i++
	}

	if hasPrice || hasAmount {
		var newPrice, newAmount any = nil, nil
		if v, ok := data["price"]; ok {
			newPrice = v
		}
		if v, ok := data["amount"]; ok {
			newAmount = v
		}

		setParts = append(setParts, fmt.Sprintf(`
			total_value = 
				COALESCE($%d, (SELECT price FROM products WHERE id = $%d)) * 
				COALESCE($%d, (SELECT amount FROM products WHERE id = $%d))
		`, i, i+2, i+1, i+2))

		values = append(values, newPrice)
		values = append(values, newAmount)
		i += 2
	}

	values = append(values, id)

	query := fmt.Sprintf(`
		UPDATE products
		SET %s
		WHERE id = $%d AND deleted_at IS NULL
		RETURNING id, created_at, updated_at, deleted_at, amount, price, total_value, name, type, ean_code, available
	`, strings.Join(setParts, ", "), i)

	err = pr.connection.QueryRow(query, values...).Scan(
		&p.ID,
		&p.CreatedAt,
		&p.UpdatedAt,
		&p.DeletedAt,
		&p.Amount,
		&p.Price,
		&p.TotalValue,
		&p.Name,
		&p.Type,
		&p.EANCode,
		&p.Available,
	)
	if err != nil {
		if err == sql.ErrNoRows {
			pr.logger.Err("produto não encontrado")
		} else {
			pr.logger.Errf("erro ao encontrar o produto: %v", err)
		}
		return models.Product{}, err
	}

	return p, nil
}
