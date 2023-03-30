package memory

import (
	"fmt"
	"golang-learn-ddd/aggregate"
	"golang-learn-ddd/domain/customer"
	"sync"

	"github.com/google/uuid"
)

type memoryRepository struct {
	customers map[uuid.UUID]aggregate.Customer
	sync.Mutex
}

func New() customer.CustomerRepository {
	return &memoryRepository{
		customers: map[uuid.UUID]aggregate.Customer{},
	}
}

func (r *memoryRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	if customer, ok := r.customers[id]; ok {
		return customer, nil
	}

	return aggregate.Customer{}, customer.ErrCustomerNotFound
}

func (r *memoryRepository) Add(c aggregate.Customer) error {
	// make sure customer is already in repository
	if _, ok := r.customers[c.GetID()]; ok {
		return fmt.Errorf("customer already exists :%w", customer.ErrFailedToAddCustomer)
	}

	// add customer to customer map
	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()

	return nil
}

func (r *memoryRepository) Update(c aggregate.Customer) error {
	if _, ok := r.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists :%w", customer.ErrUpdateCustomer)
	}

	// overwrite customer
	r.Lock()
	r.customers[c.GetID()] = c
	r.Unlock()

	return nil
}

func (r *memoryRepository) Delete(c aggregate.Customer) error {
	if _, ok := r.customers[c.GetID()]; !ok {
		return fmt.Errorf("customer does not exists :%w", customer.ErrDeleteCustomer)
	}

	// delete customer
	r.Lock()
	delete(r.customers, c.GetID())
	r.Unlock()

	return nil
}
