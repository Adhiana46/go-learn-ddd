package mongo

import (
	"context"
	"golang-learn-ddd/aggregate"
	"golang-learn-ddd/domain/customer"
	"time"

	"github.com/google/uuid"
	"github.com/pkg/errors"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type mongoRepository struct {
	db       *mongo.Database
	customer *mongo.Collection
}

// mongoCustomer internal type to store CustomerAggregate to mongodb
type mongoCustomer struct {
	ID   uuid.UUID `bson:"id"`
	Name string    `bson:"name"`
}

func NewFromCustomer(c aggregate.Customer) mongoCustomer {
	return mongoCustomer{
		ID:   c.GetID(),
		Name: c.GetName(),
	}
}

func (m *mongoCustomer) ToAggregate() aggregate.Customer {
	c := aggregate.Customer{}

	c.SetID(m.ID)
	c.SetName(m.Name)

	return c
}

func New(ctx context.Context, connectionString string) (customer.CustomerRepository, error) {
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(connectionString))
	if err != nil {
		return nil, err
	}

	db := client.Database("learn-golang-ddd")
	collection := db.Collection("customers")

	return &mongoRepository{
		db:       db,
		customer: collection,
	}, nil
}

func (r *mongoRepository) Get(id uuid.UUID) (aggregate.Customer, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var row mongoCustomer
	err := r.customer.FindOne(ctx, bson.M{"id": id}).Decode(&row)
	if err != nil {
		return aggregate.Customer{}, errors.Errorf("customer does not exists: %w; %w", err, customer.ErrCustomerNotFound)
	}

	return row.ToAggregate(), nil
}

func (r *mongoRepository) Add(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := r.customer.InsertOne(ctx, NewFromCustomer(c))
	if err != nil {
		return errors.Errorf("failed to add a customer: %w; %w", err, customer.ErrFailedToAddCustomer)
	}

	return nil
}

func (r *mongoRepository) Update(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": c.GetID()}
	updateData := bson.M{
		"$set": bson.M{
			"name": c.GetName(),
		},
	}

	_, err := r.customer.UpdateOne(ctx, filter, updateData)
	if err != nil {
		return errors.Errorf("failed to update a customer: %w; %w", err, customer.ErrUpdateCustomer)
	}

	return nil
}

func (r *mongoRepository) Delete(c aggregate.Customer) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	filter := bson.M{"id": c.GetID()}

	_, err := r.customer.DeleteOne(ctx, filter)
	if err != nil {
		return errors.Errorf("failed to delete a customer: %w; %w", err, customer.ErrDeleteCustomer)
	}

	return nil
}
