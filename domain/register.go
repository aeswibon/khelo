package domain

import (
	"context"
)

// RegisterRequest struct defining the request body
type RegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Email    string `form:"email" binding:"required,email"`
	Password string `form:"password" binding:"required"`
}

// RegisterResponse struct defining the response body
type RegisterResponse struct {
	Message string `json:"message"`
}

// RegisterUsecase interface defining methods
type RegisterUsecase interface {
	Create(c context.Context, user *User) error
	CheckUser(c context.Context, username string, email string) (bool, error)
}
