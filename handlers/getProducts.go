package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Luiz0311/estoque-go/models"
	"github.com/gin-gonic/gin"
)

func GetProducts(c *gin.Context) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, amount, price, total_value, name, type, ean_code, available 
		FROM products 
		WHERE deleted_at IS NULL
	`

	rows, err := db.Query(query)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos: " + err.Error()})
		return
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
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produto: " + err.Error()})
			return
		}

		if deletedAt.Valid {
			p.DeletedAt = &deletedAt.Time
		}

		products = append(products, p)
	}

	if !found {
		c.JSON(http.StatusOK, gin.H{"message": "Nenhum produto encontrado"})
		return
	}

	c.JSON(http.StatusOK, products)
}
