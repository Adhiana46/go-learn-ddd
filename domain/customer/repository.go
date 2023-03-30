package customer

import (
	"errors"
	"golang-learn-ddd/aggregate"

	"github.com/google/uuid"
)

var (
	ErrCustomerNotFound    = errors.New("customer not found in repository")
	ErrFailedToAddCustomer = errors.New("failed to add the customer")
	ErrUpdateCustomer      = errors.New("failed to update the customer")
	ErrDeleteCustomer      = errors.New("failed to delete the customer")
)

type CustomerRepository interface {
	Get(uuid.UUID) (aggregate.Customer, error)
	Add(aggregate.Customer) error
	Update(aggregate.Customer) error
	Delete(aggregate.Customer) error
}
