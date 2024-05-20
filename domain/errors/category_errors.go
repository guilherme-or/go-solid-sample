package errors

import "fmt"

type (
	CategoryError interface {
		error
	}

	CategoryNotFoundError struct {
		Id int64
	}

	CategoryValidationError struct {
		Field string
		Value any
	}
)

func (nfe *CategoryNotFoundError) Error() string {
	return fmt.Sprintf("Category with id %d not found", nfe.Id)
}

func (ve *CategoryValidationError) Error() string {
	return fmt.Sprintf(`Category validation failed on attribute %s. Value: "%s"`, ve.Field, ve.Value)
}
