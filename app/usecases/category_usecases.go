package usecases

import (
	"github.com/guilherme-or/go-solid-sample/app/services"
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
)

type CategoryUseCases struct {
	CategoryService services.CategoryService
}

func (cuc *CategoryUseCases) GetAllCategories() []entities.Category {
	return cuc.CategoryService.GetAll()
}

func (cuc *CategoryUseCases) GetCategoryById(id int64) (entities.Category, errors.CategoryError) {
	return cuc.CategoryService.GetById(id)
}

func (cuc *CategoryUseCases) SaveCategory(c entities.Category) (entities.Category, errors.CategoryError) {
	return cuc.CategoryService.Save(c)
}

func (cuc *CategoryUseCases) AlterCategory(c entities.Category) errors.CategoryError {
	return cuc.CategoryService.Alter(c)
}

func (cuc *CategoryUseCases) RemoveCategoryById(id int64) errors.CategoryError {
	return cuc.CategoryService.RemoveById(id)
}
