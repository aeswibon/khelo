package repository

import (
	"context"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type locationRepository struct {
	database   mongo.Database
	collection string
}

// LocationRepository function creating new repository to let user perform location related operations
func LocationRepository(db mongo.Database, collection string) domain.LocationRepository {
	return &locationRepository{
		database:   db,
		collection: collection,
	}
}

func (lr *locationRepository) FetchStates(c context.Context) ([]domain.State, error) {
	collection := lr.database.Collection(lr.collection)
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}

	var states []domain.State
	err = cursor.All(c, &states)
	if states == nil {
		return []domain.State{}, err
	}

	return states, err
}

func (lr *locationRepository) FetchDistricts(c context.Context, stateID string) ([]domain.District, error) {
	collection := lr.database.Collection(lr.collection)
	idHex, err := primitive.ObjectIDFromHex(stateID)

	var cursor mongo.Cursor
	cursor, err = collection.Find(c, bson.M{"state_id": idHex})
	if err != nil {
		return nil, err
	}

	var districts []domain.District
	err = cursor.All(c, &districts)
	if districts == nil {
		return []domain.District{}, err
	}

	return districts, err
}
