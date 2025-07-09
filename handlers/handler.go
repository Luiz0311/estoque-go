package handlers

import (
	"database/sql"
	"fmt"
	"net/http"
	"strings"
	"time"

	"github.com/Luiz0311/estoque-go/config"
	"github.com/Luiz0311/estoque-go/models"
	"github.com/Luiz0311/estoque-go/utils"
	"github.com/gin-gonic/gin"
)

var db *sql.DB

func InitializeHandler() {
	db = config.GetPostgres()
}

func GetProducts(c *gin.Context) {
	query := `
		SELECT id, created_at, updated_at, deleted_at, amount, price, name, type, ean_code, available 
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

func GetProduct(c *gin.Context) {
	id := c.Param("id")

	query := `
		SELECT id, created_at, updated_at, deleted_at, amount, price, name, type, ean_code, available
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

func CreateProdct(c *gin.Context) {
	var p models.Product

	if err := c.ShouldBind(&p); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Dados inválidos: " + err.Error()})
		return
	}

	query := `
		INSERT INTO products (name, type, ean_code, amount, price, available, created_at, updated_at)
		VALUES ($1, $2, $3, $4, $5, $6, $7, $8)
		RETURNING id, created_at, updated_at
	`
	p.CreatedAt = time.Now()
	p.UpdatedAt = time.Now()
	p.EANCode = utils.GenerateEAN13()

	err := db.QueryRow(
		query,
		p.Name,
		p.Type,
		p.EANCode,
		p.Amount,
		p.Price,
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

func UpdateProduct(c *gin.Context) {
	id := c.Param("id")

	var data map[string]interface{}
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
	values := []interface{}{}
	i := 1
	for k, v := range data {
		setParts = append(setParts, fmt.Sprintf("%s = $%d", k, i))
		values = append(values, v)
		i++
	}

	query := fmt.Sprintf(`
		UPDATE products
		SET %s
		WHERE id = $%d
		RETURNING id, created_at, updated_at
	`, strings.Join(setParts, ", "), i)

	values = append(values, id)

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

func DeleteProduct(c *gin.Context) {
	id := c.Param("id")

	query := `
		UPDATE products
		SET deleted_at = $1, updated_at = $2
		WHERE id = $3 AND deleted_at IS NULL
		RETURNING id, name
	`

	var deletedID int
	var deletedName string
	deletedAt := time.Now()
	updatedAt := time.Now()

	err := db.QueryRow(query, deletedAt, updatedAt, id).Scan(&deletedID, &deletedName)
	if err == sql.ErrNoRows {
		c.JSON(http.StatusNotFound, gin.H{"error": "produto não encontrado ou já deletado"})
		return
	} else if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "erro ao deletar produto: " + err.Error()})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"message": "produto deletado com sucesso",
		"id":      deletedID,
		"name":    deletedName,
	})
}
