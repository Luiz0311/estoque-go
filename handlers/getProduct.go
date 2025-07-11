package handlers

import (
	"database/sql"
	"net/http"

	"github.com/Luiz0311/estoque-go/models"
	"github.com/gin-gonic/gin"
)

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	query := `
		SELECT id, created_at, updated_at, deleted_at, amount, price, total_value, name, type, ean_code, available
		FROM products
		WHERE id = $1 AND deleted_at IS NULL
	`

	var p models.Product
	var deletedAt sql.NullTime

	err := db.QueryRow(query, id).Scan(
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
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao buscar o produto: " + err.Error()})
	}

	if deletedAt.Valid {
		p.DeletedAt = &deletedAt.Time
	}

	c.JSON(http.StatusOK, p)
}
