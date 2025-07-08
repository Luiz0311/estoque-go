package main

import (
	"database/sql"
	"net/http"

	"github.com/Luiz0311/estoque-go/db"
	"github.com/Luiz0311/estoque-go/models"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	db, err := db.GetDB()
	r.GET("/ping", func(ctx *gin.Context) {
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		if err = db.Ping(); err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})

	r.GET("/products", func(c *gin.Context) {
		query := `SELECT id, created_at, updated_at, deleted_at, amount, price, name, type, ean_code, available 
		          FROM products WHERE deleted_at IS NULL`

		rows, err := db.Query(query)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao buscar produtos"})
			return
		}
		defer rows.Close()

		var products []models.Product

		for rows.Next() {
			var p models.Product
			var deletedAt sql.NullTime

			err := rows.Scan(
				&p.ID,
				&p.CreatedAt,
				&p.UpdatedAt,
				&deletedAt,
				&p.Amount,
				&p.Price,
				&p.Name,
				&p.Type,
				&p.EANCode,
				&p.Available,
			)

			if err != nil {
				c.JSON(http.StatusInternalServerError, gin.H{"error": "Erro ao ler produto"})
				return
			}

			if deletedAt.Valid {
				p.DeletedAt = &deletedAt.Time
			} else {
				p.DeletedAt = nil
			}

			products = append(products, p)
		}

		c.JSON(http.StatusOK, products)
	})
	r.Run(":8080")
}
