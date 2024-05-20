package repositories

import (
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
)

type CategoryRepository interface {
	ReadAll() []entities.Category
	ReadById(id int64) (entities.Category, errors.CategoryError)
	Create(category entities.Category) (entities.Category, errors.CategoryError)
	Update(category entities.Category) errors.CategoryError
	DeleteById(id int64) errors.CategoryError
}
