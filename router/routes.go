package router

import (
	"github.com/Luiz0311/estoque-go/handlers"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api"
	handlers.InitializeHandler()

	api := r.Group(basePath)
	{
		api.GET("/products", handlers.GetProducts)
		api.GET("/product/:id", handlers.GetProduct)
		api.PATCH("/product/:id", handlers.UpdateProduct)
		api.DELETE("product/:id", handlers.DeleteProduct)
		api.POST("/product", handlers.CreateProdct)

	}
}
