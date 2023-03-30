package product

import (
	"errors"
	"golang-learn-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrProductNotFound    = errors.New("product not found in repository")
	ErrFailedToAddProduct = errors.New("failed to add the product")
	ErrUpdateProduct      = errors.New("failed to update the product")
	ErrDeleteProduct      = errors.New("failed to delete the product")
)

type ProductRepository interface {
	GetAll() ([]aggregate.Product, error)
	GetByID(id uuid.UUID) (aggregate.Product, error)
	Add(product aggregate.Product) error
	Update(product aggregate.Product) error
	Delete(id uuid.UUID) error
}
