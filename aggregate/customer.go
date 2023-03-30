package aggregate

import (
	"errors"
	"golang-learn-ddd/entity"
	"golang-learn-ddd/valueobject"

	"github.com/google/uuid"
)

var (
	ErrInvalidPerson = errors.New("a customer has to have a valid name")
)

type Customer struct {
	person   *entity.Person
	products []*entity.Item

	transaction []valueobject.Transaction
}

func NewCustomer(name string) (Customer, error) {
	if name == "" {
		return Customer{}, ErrInvalidPerson
	}

	person := &entity.Person{
		ID:   uuid.New(),
		Name: name,
	}

	return Customer{
		person:      person,
		products:    make([]*entity.Item, 0),
		transaction: make([]valueobject.Transaction, 0),
	}, nil
}

func (c *Customer) GetID() uuid.UUID {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	return c.person.ID
}

func (c *Customer) SetID(id uuid.UUID) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.ID = id
}

func (c *Customer) GetName() string {
	if c.person == nil {
		c.person = &entity.Person{}
	}
	return c.person.Name
}

func (c *Customer) SetName(name string) {
	if c.person == nil {
		c.person = &entity.Person{}
	}

	c.person.Name = name
}
