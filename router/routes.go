package router

import (
	"github.com/Luiz0311/estoque-go/config"
	"github.com/Luiz0311/estoque-go/controller"
	"github.com/Luiz0311/estoque-go/repository"
	"github.com/Luiz0311/estoque-go/usecase"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(r *gin.Engine) {
	basePath := "/api"

	dbConnection := config.GetPostgres()
	logger := config.GetLogger("repo")
	ProductRepository := repository.NewProductRepository(dbConnection, logger)
	ProductUseCase := usecase.NewProductUseCase(ProductRepository)
	ProductController := controller.NewProductController(ProductUseCase)

	api := r.Group(basePath)
	{
		api.GET("/products", ProductController.GetProducts)
		api.GET("/product/:id", ProductController.GetProduct)
		api.PATCH("/product/:id", ProductController.UpdateProduct)
		api.DELETE("product/:id", ProductController.DeleteProduct)
		api.POST("/product", ProductController.CreateProdct)

	}
}
