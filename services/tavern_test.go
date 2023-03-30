package services

import (
	"golang-learn-ddd/aggregate"
	"testing"

	"github.com/google/uuid"
)

func Test_TavernService(t *testing.T) {
	products := init_products(t)

	os, err := NewOrderService(
		WithMemoryCustomerRepository(),
		WithMemoryProductRepository(products),
	)
	if err != nil {
		t.Error(err)
	}

	tavern, err := NewTavernService(WithOrderService(os))
	if err != nil {
		t.Error(err)
	}

	cust, err := aggregate.NewCustomer("SeeU")
	if err != nil {
		t.Error(err)
	}

	if err := os.customerRepo.Add(cust); err != nil {
		t.Error(err)
	}

	order := []uuid.UUID{
		products[0].GetID(),
	}

	if err := tavern.Order(cust.GetID(), order); err != nil {
		t.Error(err)
	}
}
