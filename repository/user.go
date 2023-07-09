package repository

import (
	"context"
	"errors"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/mongo"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type userRepository struct {
	database   mongo.Database
	collection string
}

// UserRepository function creating new repository to let user perform user related operations
func UserRepository(db mongo.Database, collection string) domain.UserRepository {
	return &userRepository{
		database:   db,
		collection: collection,
	}
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	_, err := ur.GetUserByUsername(c, user.Username)
	if err == nil {
		return errors.New("username already exists")
	}
	_, err = ur.GetUserByEmail(c, user.Email)
	if err == nil {
		return errors.New("email already exists")
	}
	_, err = collection.InsertOne(c, user)
	return err
}

func (ur *userRepository) Fetch(c context.Context) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	opts := options.Find().SetProjection(bson.D{{Key: "password", Value: 0}})
	cursor, err := collection.Find(c, bson.D{}, opts)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	err = cursor.All(c, &users)
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
}

func (ur *userRepository) GetUserByUsername(c context.Context, username string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"username": username}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetUserByEmail(c context.Context, email string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	err := collection.FindOne(c, bson.M{"email": email}).Decode(&user)
	return user, err
}

func (ur *userRepository) GetByID(c context.Context, id string) (domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	var user domain.User
	idHex, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return user, err
	}
	err = collection.FindOne(c, bson.D{{Key: "_id", Value: idHex}}).Decode(&user)
	return user, err
}

func (ur *userRepository) Update(c context.Context, userID string, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	idHex, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return err
	}

	// FIXME: This is not the right way to do this
	updateFields := bson.M{
		"name":   user.Name,
		"phone":  user.Phone,
		"age":    user.Age,
		"gender": user.Gender,
	}

	// Update only those fields which are not empty
	for key, value := range updateFields {
		if str, ok := value.(string); ok && str == "" {
			delete(updateFields, key)
		}
	}

	_, err = collection.UpdateOne(c, bson.M{"_id": idHex}, bson.M{"$set": updateFields})
	return err
}
