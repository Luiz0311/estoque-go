package handlers

import (
	"net/http"
	"time"

	"github.com/Luiz0311/estoque-go/models"
	"github.com/Luiz0311/estoque-go/utils"
	"github.com/gin-gonic/gin"
)

func CreateProdct(c *gin.Context) {
	var p models.Product

	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

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

	err := db.QueryRow(
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
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao criar o produto: " + err.Error()})
		return
	}

	c.JSON(http.StatusCreated, p)
}
