package domain

import (
	"context"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

const (
	// CollectionFacility name of the collection storing facility related documents
	CollectionFacility = "facilities"
)

// Facility struct defining fields in facility model
type Facility struct {
	ID                 primitive.ObjectID `bson:"_id"`
	Name               string             `bson:"name" json:"name"`
	Email              string             `bson:"email" json:"email"`
	Phone              string             `bson:"phone" json:"phone"`
	RegistrationNumber string             `bson:"registration_number" json:"registration_number"`
	Ownership          int                `bson:"ownership" json:"ownership"`
	Active             bool               `bson:"active" json:"active"`
	Verified           bool               `bson:"verified" json:"verified"`
	Deleted            bool               `bson:"deleted" json:"deleted"`
	CreatedBy          string             `bson:"created_by" json:"created_by"`
}

// FacilityRepository interface defining database operations on facility model
type FacilityRepository interface {
	Create(c context.Context, facility *Facility) error
	Fetch(c context.Context) ([]Facility, error)
	GetFacilityByName(c context.Context, name string) (Facility, error)
	GetFacilityByEmail(c context.Context, email string) (Facility, error)
	GetByID(c context.Context, id string) (Facility, error)
	Update(c context.Context, facilityID string, facility *Facility) error
	DeleteByID(c context.Context, id string) error
}

// FacilityUsecase interface defining business logic operations on facility model
type FacilityUsecase interface {
	Create(c context.Context, facility *Facility) error
	Fetch(c context.Context) ([]Facility, error)
	GetFacilityByName(c context.Context, name string) (Facility, error)
	GetFacilityByEmail(c context.Context, email string) (Facility, error)
	GetFacilityByID(c context.Context, id string) (Facility, error)
	UpdateFacility(c context.Context, facilityID string, facility *Facility) error
	// DeleteByID(c context.Context, id string) error
}
