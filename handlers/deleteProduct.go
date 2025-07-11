package handlers

import (
	"database/sql"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

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
