package services

import (
	"context"
	"fmt"
	"golang-learn-ddd/aggregate"
	"golang-learn-ddd/domain/customer"
	customerMemory "golang-learn-ddd/domain/customer/memory"
	customerMongo "golang-learn-ddd/domain/customer/mongo"
	"golang-learn-ddd/domain/product"
	productMemory "golang-learn-ddd/domain/product/memory"

	"github.com/google/uuid"
)

type OrderConfiguration func(os *OrderService) error

type OrderService struct {
	customerRepo customer.CustomerRepository
	productRepo  product.ProductRepository
}

func NewOrderService(cfgs ...OrderConfiguration) (*OrderService, error) {
	os := &OrderService{}

	// Loop through all cfgs and apply them
	for _, cfg := range cfgs {
		err := cfg(os)

		if err != nil {
			return nil, err
		}
	}

	return os, nil
}

func WithMemoryCustomerRepository() OrderConfiguration {
	repo := customerMemory.New()
	return WithCustomerRepository(repo)
}

func WithMongoCustomerRepository(ctx context.Context, connectionString string) OrderConfiguration {
	return func(os *OrderService) error {
		repo, err := customerMongo.New(ctx, connectionString)
		if err != nil {
			return err
		}

		os.customerRepo = repo

		return nil
	}
}

func WithCustomerRepository(customerRepo customer.CustomerRepository) OrderConfiguration {
	return func(os *OrderService) error {
		os.customerRepo = customerRepo
		return nil
	}
}

func WithMemoryProductRepository(products []aggregate.Product) OrderConfiguration {
	return func(os *OrderService) error {
		repo := productMemory.New()

		for _, p := range products {
			if err := repo.Add(p); err != nil {
				return err
			}
		}

		os.productRepo = repo
		return nil
	}
}

func (os *OrderService) CreateOrder(customerID uuid.UUID, productsIDs []uuid.UUID) (float64, error) {
	// Fetch the customer
	c, err := os.customerRepo.Get(customerID)
	if err != nil {
		return 0, err
	}

	// Get each product
	var products []aggregate.Product
	var total float64

	for _, id := range productsIDs {
		p, err := os.productRepo.GetByID(id)
		if err != nil {
			return 0, err
		}

		products = append(products, p)
		total += p.GetPrice()
	}

	fmt.Printf("Customer: %s has ordered %d products with total of $%.2f\n", c.GetID(), len(products), total)
	for _, p := range products {
		fmt.Printf("\t - %s\n", p.GetItem().Name)
	}

	return total, nil
}
