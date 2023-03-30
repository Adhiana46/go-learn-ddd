package memory

import (
	"fmt"
	"golang-learn-ddd/aggregate"
	"golang-learn-ddd/domain/product"
	"sync"

	"github.com/google/uuid"
)

type memoryRepository struct {
	products map[uuid.UUID]aggregate.Product
	sync.Mutex
}

func New() product.ProductRepository {
	return &memoryRepository{
		products: map[uuid.UUID]aggregate.Product{},
	}
}

func (r *memoryRepository) GetAll() ([]aggregate.Product, error) {
	var products []aggregate.Product

	for _, product := range r.products {
		products = append(products, product)
	}

	return products, nil
}

func (r *memoryRepository) GetByID(id uuid.UUID) (aggregate.Product, error) {
	if product, ok := r.products[id]; ok {
		return product, nil
	}

	return aggregate.Product{}, product.ErrProductNotFound
}

func (r *memoryRepository) Add(p aggregate.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[p.GetID()]; ok {
		return fmt.Errorf("product already exists :%w", product.ErrFailedToAddProduct)
	}

	r.products[p.GetID()] = p

	return nil
}

func (r *memoryRepository) Update(p aggregate.Product) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[p.GetID()]; !ok {
		return fmt.Errorf("product is not exists :%w", product.ErrUpdateProduct)
	}

	r.products[p.GetID()] = p

	return nil
}

func (r *memoryRepository) Delete(id uuid.UUID) error {
	r.Lock()
	defer r.Unlock()

	if _, ok := r.products[id]; !ok {
		return fmt.Errorf("product is not exists :%w", product.ErrDeleteProduct)
	}

	delete(r.products, id)

	return nil
}
