package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/gin-gonic/gin"
)

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var data map[string]any
	if err := c.BindJSON(&data); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "JSON inválido"})
		return
	}

	if len(data) == 0 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "nenhum campo enviado"})
		return
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
		RETURNING id, created_at, updated_at
	`, strings.Join(setParts, ", "), i)

	var updatedID int
	var createdAt, updatedAt time.Time

	err := db.QueryRow(query, values...).Scan(&updatedID, &createdAt, &updatedAt)
	if err != nil {
		if err == sql.ErrNoRows {
			c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado"})
		} else {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao atualizar o produto: " + err.Error()})
		}
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message":    "produto atualizado com sucesso",
		"id":         updatedID,
		"created_at": createdAt,
		"updated_at": updatedAt,
	})
}
