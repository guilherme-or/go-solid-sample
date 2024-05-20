package services

import (
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
	"github.com/guilherme-or/go-solid-sample/domain/repositories"
)

type ProductService struct {
	ProductRepo repositories.ProductRepository
}

func (cs *ProductService) GetAll() []entities.Product {
	return cs.ProductRepo.ReadAll()
}

func (cs *ProductService) GetById(id int64) (entities.Product, errors.ProductError) {
	return cs.ProductRepo.ReadById(id)
}

func (cs *ProductService) Save(c entities.Product) (entities.Product, errors.ProductError) {
	return cs.ProductRepo.Create(c)
}

func (cs *ProductService) Alter(c entities.Product) errors.ProductError {
	return cs.ProductRepo.Update(c)
}

func (cs *ProductService) RemoveById(id int64) errors.ProductError {
	return cs.ProductRepo.DeleteById(id)
}
