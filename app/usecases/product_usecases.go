package usecases

import (
	"github.com/guilherme-or/go-solid-sample/app/services"
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
)

type ProductUseCases struct {
	ProductService services.ProductService
}

func (puc *ProductUseCases) GetAllProducts() []entities.Product {
	return puc.ProductService.GetAll()
}

func (puc *ProductUseCases) GetProductById(id int64) (entities.Product, errors.ProductError) {
	return puc.ProductService.GetById(id)
}

func (puc *ProductUseCases) SaveProduct(p entities.Product) (entities.Product, errors.ProductError) {
	return puc.ProductService.Save(p)
}

func (puc *ProductUseCases) AlterProduct(p entities.Product) errors.ProductError {
	return puc.ProductService.Alter(p)
}

func (puc *ProductUseCases) RemoveProductById(id int64) errors.ProductError {
	return puc.ProductService.RemoveById(id)
}
