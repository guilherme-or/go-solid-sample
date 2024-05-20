package errors

import "fmt"

type (
	ProductError interface {
		error
	}

	ProductNotFoundError struct {
		Id int64
	}

	ProductValidationError struct {
		Field string
		Value any
	}
)

func (nfe *ProductNotFoundError) Error() string {
	return fmt.Sprintf("Product with id %d not found", nfe.Id)
}

func (ve *ProductValidationError) Error() string {
	return fmt.Sprintf(`Product validation failed on field %s. Value: "%s"`, ve.Field, ve.Value)
}
