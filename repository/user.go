package repository

import (
	"context"
	"errors"

	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/internal"
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

func checkUnique(c context.Context, collection mongo.Collection, field string, value string) bool {
	user, err := collection.CountDocuments(c, bson.M{field: value})
	if err != nil {
		return false
	}
	return user == 0
}

func (ur *userRepository) Create(c context.Context, user *domain.User) error {
	collection := ur.database.Collection(ur.collection)
	if check := checkUnique(c, collection, "username", user.Username); check {
		return errors.New("username already exists")
	}
	if check := checkUnique(c, collection, "email", user.Email); check {
		return errors.New("email already exists")
	}

	// Hash password before storing in database
	hashPassword, err := internal.HashPassword(user.Password)
	if err != nil {
		return errors.New("Internal server error")
	}
	user.Password = hashPassword
	_, err = collection.InsertOne(c, user)
	return err
}

func (ur *userRepository) Fetch(c context.Context, filter interface{}, projection interface{}) ([]domain.User, error) {
	collection := ur.database.Collection(ur.collection)
	opts := options.Find().SetProjection(projection)
	cursor, err := collection.Find(c, filter, opts)
	if err != nil {
		return nil, err
	}

	var users []domain.User
	for cursor.Next(c) {
		var user domain.User
		err := cursor.Decode(&user)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	if users == nil {
		return []domain.User{}, err
	}

	return users, err
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
