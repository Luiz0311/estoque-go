package models

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SendoError(ctx *gin.Context, status int, msg string) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(status, gin.H{
		"error": msg,
	})
}

func SendSuccess(ctx *gin.Context, data any) {
	ctx.Header("Content-Type", "application/json")
	ctx.JSON(http.StatusOK, gin.H{
		"data": data,
	})
}
