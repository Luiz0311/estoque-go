package controller

import (
	"net/http"

	m "github.com/Luiz0311/estoque-go/models"
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
	if err == m.ErrDB || err == m.ErrRead {
		m.SendoError(ctx, http.StatusInternalServerError, err.Error())
		return
	} else if err != nil {
		m.SendoError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	m.SendSuccess(ctx, products)
}

func (pc *ProductController) GetProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := pc.productUseCase.GetProduct(id)
	if err != nil {
		m.SendoError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	m.SendSuccess(ctx, product)
}

func (pc *ProductController) DeleteProduct(ctx *gin.Context) {
	id := ctx.Param("id")

	product, err := pc.productUseCase.GetProduct(id)
	if err != nil {
		m.SendoError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	m.SendSuccess(ctx, product)
}

func (pc *ProductController) CreateProdct(ctx *gin.Context) {
	var p m.Product

	if err := ctx.ShouldBind(&p); err != nil {
		m.SendoError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product, err := pc.productUseCase.CreateProdct(p)
	if err != nil {
		m.SendoError(ctx, http.StatusInternalServerError, err.Error())
		return
	}

	m.SendSuccess(ctx, product)
}

func (pc ProductController) UpdateProduct(ctx *gin.Context) {
	id := ctx.Param("id")
	var data map[string]any
	if err := ctx.BindJSON(&data); err != nil {
		m.SendoError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	product, err := pc.productUseCase.UpdateProduct(id, data)
	if err != nil {
		m.SendoError(ctx, http.StatusBadRequest, err.Error())
		return
	}

	m.SendSuccess(ctx, product)
}
