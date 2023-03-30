package memory

import (
	"errors"
	"golang-learn-ddd/aggregate"
	"golang-learn-ddd/domain/customer"
	"testing"

	"github.com/google/uuid"
)

func Test_memoryRepository_Get(t *testing.T) {
	repo := New()

	cust, err := aggregate.NewCustomer("Adhiana")
	if err != nil {
		t.Fatal(err)
	}

	repo.Add(cust)

	type testCase struct {
		name        string
		id          uuid.UUID
		expectedErr error
	}
	tests := []testCase{
		{
			name:        "customer not found",
			id:          uuid.New(),
			expectedErr: customer.ErrCustomerNotFound,
		},
		{
			name:        "customer found",
			id:          cust.GetID(),
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := repo.Get(tt.id)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func Test_memoryRepository_Add(t *testing.T) {
	repo := New()

	cust, err := aggregate.NewCustomer("Adhiana")
	if err != nil {
		t.Fatal(err)
	}

	type testCase struct {
		name        string
		cust        aggregate.Customer
		expectedErr error
	}
	tests := []testCase{
		{
			name:        "add new customer",
			cust:        cust,
			expectedErr: nil,
		},
		{
			name:        "add custumer which is already exists",
			cust:        cust,
			expectedErr: customer.ErrFailedToAddCustomer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Add(tt.cust)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func Test_memoryRepository_Update(t *testing.T) {
	repo := New()

	cust1, err := aggregate.NewCustomer("Adhiana")
	if err != nil {
		t.Fatal(err)
	}
	cust2, err := aggregate.NewCustomer("Mastur")
	if err != nil {
		t.Fatal(err)
	}

	// add only cust1 to the repo
	repo.Add(cust1)

	type testCase struct {
		name        string
		cust        aggregate.Customer
		expectedErr error
	}
	tests := []testCase{
		{
			name:        "update existing customer",
			cust:        cust1,
			expectedErr: nil,
		},
		{
			name:        "update customer that does not exists",
			cust:        cust2,
			expectedErr: customer.ErrUpdateCustomer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Update(tt.cust)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func Test_memoryRepository_Delete(t *testing.T) {
	repo := New()

	cust1, err := aggregate.NewCustomer("Adhiana")
	if err != nil {
		t.Fatal(err)
	}
	cust2, err := aggregate.NewCustomer("Mastur")
	if err != nil {
		t.Fatal(err)
	}

	// add only cust1 to the repo
	repo.Add(cust1)

	type testCase struct {
		name        string
		cust        aggregate.Customer
		expectedErr error
	}
	tests := []testCase{
		{
			name:        "delete existing customer",
			cust:        cust1,
			expectedErr: nil,
		},
		{
			name:        "delete customer that does not exists",
			cust:        cust2,
			expectedErr: customer.ErrDeleteCustomer,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			err := repo.Delete(tt.cust)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}
