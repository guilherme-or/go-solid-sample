package services

import (
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
	"github.com/guilherme-or/go-solid-sample/domain/repositories"
)

type CategoryService struct {
	CategoryRepo repositories.CategoryRepository
}

func (cs *CategoryService) GetAll() []entities.Category {
	return cs.CategoryRepo.ReadAll()
}

func (cs *CategoryService) GetById(id int64) (entities.Category, errors.CategoryError) {
	return cs.CategoryRepo.ReadById(id)
}

func (cs *CategoryService) Save(c entities.Category) (entities.Category, errors.CategoryError) {
	return cs.CategoryRepo.Create(c)
}

func (cs *CategoryService) Alter(c entities.Category) errors.CategoryError {
	return cs.CategoryRepo.Update(c)
}

func (cs *CategoryService) RemoveById(id int64) errors.CategoryError {
	return cs.CategoryRepo.DeleteById(id)
}
