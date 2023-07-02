package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionLocation name of the collection storing location related documents
	CollectionLocation = "locations"
)

// State struct defining fields in state model
type State struct {
	ID   primitive.ObjectID `bson:"id"`
	Name string             `bson:"name"`
}

// District struct defining fields in district model
type District struct {
	ID      primitive.ObjectID `bson:"id"`
	Name    string             `bson:"name"`
	StateID primitive.ObjectID `bson:"state_id"`
}

// LocationRepository interface defining database operations on location model
type LocationRepository interface {
	FetchStates(c context.Context) ([]State, error)
	FetchDistricts(c context.Context, stateID string) ([]District, error)
}
