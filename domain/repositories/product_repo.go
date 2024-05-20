package repositories

import (
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
)

type ProductRepository interface {
	ReadAll() []entities.Product
	ReadById(id int64) (entities.Product, errors.ProductError)
	Create(product entities.Product) (entities.Product, errors.ProductError)
	Update(product entities.Product) errors.ProductError
	DeleteById(id int64) errors.ProductError
}
