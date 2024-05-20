package persistance

import (
	"github.com/guilherme-or/go-solid-sample/domain/entities"
	"github.com/guilherme-or/go-solid-sample/domain/errors"
	"reflect"
)

type LocalProductRepository struct {
	products []entities.Product
}

func NewLocalProductRepository() *LocalProductRepository {
	return &LocalProductRepository{
		[]entities.Product{
			{
				Id:          1,
				Name:        "Product No. 1",
				Description: "Product No. 1 Description",
				Price:       100.11,
				CategoryId:  1,
				Active:      true,
			},
			{
				Id:          2,
				Name:        "Product No. 2",
				Description: "Product No. 2 Description",
				Price:       200.22,
				CategoryId:  2,
				Active:      true,
			},
			{
				Id:          3,
				Name:        "Product No. 3",
				Description: "Product No. 3 Description",
				Price:       300.33,
				CategoryId:  3,
				Active:      false,
			},
			{
				Id:          4,
				Name:        "Product No. 4",
				Description: "Product No. 4 Description",
				Price:       400.44,
				CategoryId:  3,
				Active:      false,
			},
			{
				Id:          5,
				Name:        "Product No. 5",
				Description: "Product No. 5 Description",
				Price:       500.55,
				CategoryId:  2,
				Active:      true,
			},
			{
				Id:          6,
				Name:        "Product No. 6",
				Description: "Product No. 6 Description",
				Price:       600.66,
				CategoryId:  1,
				Active:      false,
			},
		},
	}
}

func (lpr *LocalProductRepository) ReadAll() []entities.Product {
	return lpr.products
}

func (lpr *LocalProductRepository) ReadById(id int64) (entities.Product, errors.ProductError) {
	for _, p := range lpr.products {
		if p.Id == id {
			return p, nil
		}
	}

	return entities.Product{}, &errors.ProductNotFoundError{Id: id}
}

func (lpr *LocalProductRepository) Create(p entities.Product) (entities.Product, errors.ProductError) {
	if err := lpr.validateProduct(&p); err != nil {
		return entities.Product{}, err
	}

	lpr.products = append(lpr.products, p)
	return p, nil
}

func (lpr *LocalProductRepository) Update(p entities.Product) errors.ProductError {
	if err := lpr.validateProduct(&p); err != nil {
		return err
	}

	for i, product := range lpr.products {
		if product.Id == p.Id {
			lpr.products[i] = p
			return nil
		}
	}

	return &errors.ProductNotFoundError{Id: p.Id}
}

func (lpr *LocalProductRepository) DeleteById(id int64) errors.ProductError {
	for i, p := range lpr.products {
		if p.Id == id {
			lpr.products = append(lpr.products[:i], lpr.products[i+1:]...)
			return nil
		}
	}

	return &errors.ProductNotFoundError{Id: id}
}

func (lpr *LocalProductRepository) validateProduct(p *entities.Product) errors.ProductError {
	t := reflect.TypeOf(*p)
	v := reflect.ValueOf(*p)

	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)

		if value.IsZero() {
			return &errors.ProductValidationError{Field: field.Name, Value: value}
		}
	}

	return nil
}
