package persistance

import (
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
	"reflect"
)

type LocalCategoryRepository struct {
	categories []entities.Category
}

func NewLocalCategoryRepository() *LocalCategoryRepository {
	return &LocalCategoryRepository{
		[]entities.Category{
			{
				Id:          1,
				Name:        "Category No. 1",
				Description: "Category No. 1 Description",
			},
			{
				Id:          2,
				Name:        "Category No. 2",
				Description: "Category No. 2 Description",
			},
			{
				Id:          3,
				Name:        "Category No. 3",
				Description: "Category No. 3 Description",
			},
		},
	}
}

func (lpr *LocalCategoryRepository) ReadAll() []entities.Category {
	return lpr.categories
}

func (lpr *LocalCategoryRepository) ReadById(id int64) (entities.Category, errors.CategoryError) {
	for _, p := range lpr.categories {
		if p.Id == id {
			return p, nil
		}
	}

	return entities.Category{}, &errors.CategoryNotFoundError{Id: id}
}

func (lpr *LocalCategoryRepository) Create(p entities.Category) (entities.Category, errors.CategoryError) {
	if err := lpr.validateCategory(&p); err != nil {
		return entities.Category{}, err
	}

	lpr.categories = append(lpr.categories, p)
	return p, nil
}

func (lpr *LocalCategoryRepository) Update(p entities.Category) errors.CategoryError {
	if err := lpr.validateCategory(&p); err != nil {
		return err
	}

	for i, category := range lpr.categories {
		if category.Id == p.Id {
			lpr.categories[i] = p
			return nil
		}
	}

	return &errors.CategoryNotFoundError{Id: p.Id}
}

func (lpr *LocalCategoryRepository) DeleteById(id int64) errors.CategoryError {
	for i, p := range lpr.categories {
		if p.Id == id {
			lpr.categories = append(lpr.categories[:i], lpr.categories[i+1:]...)
			return nil
		}
	}

	return &errors.CategoryNotFoundError{Id: id}
}

func (lpr *LocalCategoryRepository) validateCategory(p *entities.Category) errors.CategoryError {
	t := reflect.TypeOf(*p)
	v := reflect.ValueOf(*p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.IsZero() {
			return &errors.CategoryValidationError{Field: field.Name, Value: value}
		}
	}

	return nil
}
