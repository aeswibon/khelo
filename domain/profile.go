package domain

import (
	"context"
)

// Profile ...
type Profile struct {
	Username string `json:"username"`
	Name     string `json:"name"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
	Gender   string `json:"gender"`
	Age      int    `json:"age"`
}

// ProfileUsecase ...
type ProfileUsecase interface {
	GetProfileByID(c context.Context, userID string) (*Profile, error)
	UpdateProfile(c context.Context, userID string, profile *User) error
}
