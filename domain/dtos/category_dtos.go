package dtos

import "github.com/guilherme-or/go-solid-sample/domain/entities"

type (
	ReadCategoryDTO struct {
		Id          int64
		Name        string
		Description string
		Products    []entities.Product
	}

	CreateCategoryDTO struct {
		Name        string
		Description string
	}
)
