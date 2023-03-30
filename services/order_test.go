package services

import (
	"golang-learn-ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func init_products(t *testing.T) []aggregate.Product {
	beer, err := aggregate.NewProduct("Beer", "Halal Beer", 99.92)
	if err != nil {
		t.Fatal(err)
	}

	peanut, err := aggregate.NewProduct("Peanut Butter", "Peanut Nut Day", 12.5)
	if err != nil {
		t.Fatal(err)
	}

	bakso, err := aggregate.NewProduct("Bakso Kuah", "Bakso Kuah pedah hot jeletot", 1.5)
	if err != nil {
		t.Fatal(err)
	}

	return []aggregate.Product{
		beer, peanut, bakso,
	}
}

func TestOrder_NewOrderService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryProductRepository(products),
		WithMemoryCustomerRepository(),
	)

	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("Senyamiku")
	if err != nil {
		t.Error(err)
	}

	if err := os.customerRepo.Add(cust); err != nil {
		t.Error(err)
	}

	orders := []uuid.UUID{
		products[0].GetID(),
		products[2].GetID(),
	}

	if _, err := os.CreateOrder(cust.GetID(), orders); err != nil {
		t.Error(err)
	}
}
