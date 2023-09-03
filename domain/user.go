package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionUser name of the collection storing user related documents
	CollectionUser = "users"
)

// User struct defining fields in user model
type User struct {
	ID       primitive.ObjectID `bson:"_id"`
	Username string             `bson:"username" json:"username"`
	Name     string             `bson:"name" json:"name,omitempty"`
	Email    string             `bson:"email" json:"email,omitempty"`
	Password string             `bson:"password" json:"-"`
	Type     int                `bson:"type" json:"type,omitempty"`
	Phone    string             `bson:"phone,omitempty" json:"phone,omitempty"`
	Gender   string             `bson:"gender" json:"gender,omitempty"`
	Age      int                `bson:"age" json:"age,omitempty"`
	Deleted  bool               `bson:"deleted"`
}

// UserRepository interface defining database operations on user model
type UserRepository interface {
	Create(c context.Context, user *User) error
	Fetch(c context.Context, filter interface{}, projection interface{}) ([]User, error)
	GetByID(c context.Context, id string) (User, error)
	Update(c context.Context, userID string, user *User) error
}
