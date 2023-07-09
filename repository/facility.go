package repository

import (
	"context"
	"errors"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type facilityRepository struct {
	database   mongo.Database
	collection string
}

// FacilityRepository function creating new repository to let user perform facility related operations
func FacilityRepository(db mongo.Database, collection string) domain.FacilityRepository {
	return &facilityRepository{
		database:   db,
		collection: collection,
	}
}

func (fr *facilityRepository) Create(c context.Context, facility *domain.
	Facility) error {
	collection := fr.database.Collection(fr.collection)
	_, err := fr.GetFacilityByName(c, facility.Name)
	if err == nil {
		return errors.New("facility name already exists")
	}
	_, err = fr.GetFacilityByEmail(c, facility.Email)
	_, err = collection.InsertOne(c, facility)
	return err
}

func (fr *facilityRepository) Fetch(c context.Context) ([]domain.Facility, error) {
	collection := fr.database.Collection(fr.collection)
	cursor, err := collection.Find(c, bson.D{})
	if err != nil {
		return nil, err
	}

	var facilities []domain.Facility
	err = cursor.All(c, &facilities)
	if facilities == nil {
		return []domain.Facility{}, err
	}
	return facilities, err
}

func (fr *facilityRepository) GetFacilityByName(c context.Context, name string) (domain.Facility, error) {
	collection := fr.database.Collection(fr.collection)
	var facility domain.Facility
	err := collection.FindOne(c, bson.D{{Key: "name", Value: name}}).Decode(&facility)
	return facility, err
}

func (fr *facilityRepository) GetFacilityByEmail(c context.Context, email string) (domain.Facility, error) {
	collection := fr.database.Collection(fr.collection)
	var facility domain.Facility
	err := collection.FindOne(c, bson.D{{Key: "email", Value: email}}).Decode(&facility)
	return facility, err
}

func (fr *facilityRepository) GetByID(c context.Context, id string) (domain.Facility, error) {
	collection := fr.database.Collection(fr.collection)
	var facility domain.Facility
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return facility, err
	}
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: idHex}}).Decode(&facility)
	return facility, err
}

func (fr *facilityRepository) Update(c context.Context, facilityID string, facility *domain.Facility) error {
	collection := fr.database.Collection(fr.collection)
	idHex, err := primitive.ObjectIDFromHex(facilityID)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(c, bson.D{{Key: "_id", Value: idHex}}, bson.D{{Key: "$set", Value: facility}})
	return err
}

func (fr *facilityRepository) DeleteByID(c context.Context, id string) error {
	collection := fr.database.Collection(fr.collection)
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}
	_, err = collection.UpdateOne(c, bson.D{{Key: "_id", Value: idHex}}, bson.D{{Key: "$set", Value: bson.D{{Key: "deleted", Value: true}}}})
	return err
}
