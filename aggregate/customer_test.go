package aggregate

import (
	"errors"
	"golang-learn-ddd/entity"
	"golang-learn-ddd/valueobject"
	"reflect"
	"testing"

	"github.com/google/uuid"
)

func TestCustomer_NewCustomer(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name        string
		args        args
		expectedErr error
	}{
		{
			name: "Empty name validation",
			args: args{
				name: "",
			},
			expectedErr: ErrInvalidPerson,
		},
		{
			name: "Valid name",
			args: args{
				name: "Adhiana",
			},
			expectedErr: nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			_, err := NewCustomer(tt.args.name)
			if !errors.Is(err, tt.expectedErr) {
				t.Errorf("expected error %v, got %v", tt.expectedErr, err)
			}
		})
	}
}

func TestCustomer_GetID(t *testing.T) {
	id1 := uuid.New()

	type fields struct {
		person      *entity.Person
		products    []*entity.Item
		transaction []valueobject.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		want   uuid.UUID
	}{
		{
			name: "Get Valid ID",
			fields: fields{
				person: &entity.Person{
					ID:   id1,
					Name: "Adhiana",
				},
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			want: id1,
		},
		{
			name: "empty person",
			fields: fields{
				person:      nil,
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			want: [16]byte{},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customer{
				person:      tt.fields.person,
				products:    tt.fields.products,
				transaction: tt.fields.transaction,
			}
			if got := c.GetID(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Customer.GetID() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_SetID(t *testing.T) {
	id := uuid.New()
	newId := uuid.New()

	type fields struct {
		person      *entity.Person
		products    []*entity.Item
		transaction []valueobject.Transaction
	}
	type args struct {
		id uuid.UUID
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set ID",
			fields: fields{
				person: &entity.Person{
					ID:   id,
					Name: "Adhiana",
				},
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			args: args{
				id: newId,
			},
		},
		{
			name: "Set ID if person is null",
			fields: fields{
				person:      nil,
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			args: args{
				id: newId,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customer{
				person:      tt.fields.person,
				products:    tt.fields.products,
				transaction: tt.fields.transaction,
			}
			c.SetID(tt.args.id)

			if c.GetID() != tt.args.id {
				t.Errorf("expecting id to be %v, but got %v", tt.args.id, c.GetID())
			}
		})
	}
}

func TestCustomer_GetName(t *testing.T) {
	type fields struct {
		person      *entity.Person
		products    []*entity.Item
		transaction []valueobject.Transaction
	}
	tests := []struct {
		name   string
		fields fields
		want   string
	}{
		{
			name: "Get Valid Name",
			fields: fields{
				person: &entity.Person{
					ID:   uuid.New(),
					Name: "Adhiana",
				},
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			want: "Adhiana",
		},
		{
			name: "empty person",
			fields: fields{
				person:      nil,
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			want: "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customer{
				person:      tt.fields.person,
				products:    tt.fields.products,
				transaction: tt.fields.transaction,
			}
			if got := c.GetName(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Customer.GetName() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestCustomer_SetName(t *testing.T) {
	name := "Adhiana"
	newName := "Mastur"

	type fields struct {
		person      *entity.Person
		products    []*entity.Item
		transaction []valueobject.Transaction
	}
	type args struct {
		name string
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		{
			name: "Set Name",
			fields: fields{
				person: &entity.Person{
					ID:   uuid.New(),
					Name: name,
				},
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			args: args{
				name: newName,
			},
		},
		{
			name: "Set Name if person is null",
			fields: fields{
				person:      nil,
				products:    []*entity.Item{},
				transaction: []valueobject.Transaction{},
			},
			args: args{
				name: newName,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			c := &Customer{
				person:      tt.fields.person,
				products:    tt.fields.products,
				transaction: tt.fields.transaction,
			}
			c.SetName(tt.args.name)

			if c.GetName() != tt.args.name {
				t.Errorf("expecting name to be %v, but got %v", tt.args.name, c.GetName())
			}
		})
	}
}
