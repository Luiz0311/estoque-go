package usecase

import (
	"strconv"

	"github.com/Luiz0311/estoque-go/models"
	"github.com/Luiz0311/estoque-go/repository"
)

type ProductUseCase struct {
	repository repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return ProductUseCase{
		repository: repo,
	}
}

func (pu *ProductUseCase) GetProducts() ([]models.Product, error) {
	return pu.repository.GetProducts()
}

func (pu *ProductUseCase) GetProduct(id string) (models.Product, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, err
	}

	product, err := pu.repository.GetProduct(intId)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pu *ProductUseCase) CreateProdct(product models.Product) (models.Product, error) {
	product, err := pu.repository.CreateProdct(product)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pu *ProductUseCase) DeleteProduct(id string) (models.Product, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, err
	}

	product, err := pu.repository.DeleteProduct(intId)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}

func (pu *ProductUseCase) UpdateProduct(id string, data map[string]any) (models.Product, error) {
	intId, err := strconv.Atoi(id)
	if err != nil {
		return models.Product{}, err
	}

	product, err := pu.repository.UpdateProduct(intId, data)
	if err != nil {
		return models.Product{}, err
	}

	return product, nil
}
