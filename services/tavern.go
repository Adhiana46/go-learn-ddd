package services

import (
	"fmt"

	"github.com/google/uuid"
)

type TavernConfiguration func(s *TavernService) error

type TavernService struct {
	OrderService *OrderService

	BillingService interface{}
}

func NewTavernService(cfgs ...TavernConfiguration) (*TavernService, error) {
	s := &TavernService{}

	for _, cfg := range cfgs {
		if err := cfg(s); err != nil {
			return nil, err
		}
	}

	return s, nil
}

func WithOrderService(os *OrderService) TavernConfiguration {
	return func(s *TavernService) error {
		s.OrderService = os
		return nil
	}
}

func (s *TavernService) Order(customer uuid.UUID, products []uuid.UUID) error {
	price, err := s.OrderService.CreateOrder(customer, products)
	if err != nil {
		return err
	}

	fmt.Printf("\nBill the customer: %.2f\n\n", price)

	return nil
}
