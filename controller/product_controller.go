package controller

import (
	"net/http"

	"github.com/Luiz0311/estoque-go/models"
	"github.com/Luiz0311/estoque-go/usecase"
	"github.com/gin-gonic/gin"
)

type ProductController struct {
	productUseCase usecase.ProductUseCase
}

func NewProductController(useCase usecase.ProductUseCase) ProductController {
	return ProductController{productUseCase: useCase}
}

func (pc *ProductController) GetProducts(ctx *gin.Context) {
	products, err := pc.productUseCase.GetProducts()
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, products)
}

func (pc *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := pc.productUseCase.GetProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := pc.productUseCase.DeleteProduct(id)
	if err != nil {
		ctx.JSON(http.StatusInternalServerError, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc *ProductController) CreateProdct(ctx *gin.Context) {
	var p models.Product

	if err := ctx.ShouldBind(&p); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := pc.productUseCase.CreateProdct(p)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}

func (pc ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var data map[string]any
	if err := ctx.BindJSON(&data); err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	product, err := pc.productUseCase.UpdateProduct(id, data)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, product)
}
