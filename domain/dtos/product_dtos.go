package dtos

import "github.com/guilherme-or/go-solid-sample/domain/entities"

type (
	ReadProductDTO struct {
		Id          int64
		Name        string
		Description string
		Price       float64
		Category    entities.Category
		Active      bool
	}

	CreateProductDTO struct {
		Name        string
		Description string
		Price       float64
		CategoryId  int64
		Active      bool
	}
)
