package domain

import (
	"context"
)

// RegisterRequest struct defining the request body
type RegisterRequest struct {
	Username string `form:"username" binding:"required"`
	Name     string `form:"name" binding:"required"`
	Email    string `form:"email" binding:"required"`
	Phone    string `form:"phone" binding:"required"`
	Gender   string `form:"gender" binding:"required"`
	Age      int    `form:"age" binding:"required"`
	Password string `form:"password" binding:"required"`
}

// RegisterResponse struct defining the response body
type RegisterResponse struct {
	Message string `json:"message"`
}

// RegisterUsecase interface defining methods
type RegisterUsecase interface {
	Create(c context.Context, user *RegisterRequest) error
}
