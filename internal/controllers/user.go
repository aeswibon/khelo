package controllers

import (
	"net/http"
	"strings"

	"github.com/cp-Coder/khelo/internal/models"
	"github.com/cp-Coder/khelo/pkg/utils"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// UserController ...
type UserController struct{}

var userModel = &models.UserModel{}

// getUserID ...
func getUserID(c *gin.Context) (userID uuid.UUID) {
	// MustGet returns the value for the given key if it exists, otherwise it panics.
	return c.MustGet("userID").(uuid.UUID)
}

// Login controller handling the login request
func (ctrl *UserController) Login(c *gin.Context) {
	data, err := utils.ParseRequestBody(c, &models.LoginForm{})
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "Invalid request"})
		return
	}

	loginForm := data.(*models.LoginForm)
	user, token, err := userModel.Login(loginForm)

	if err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": "Invalid login details"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged in", "user": user, "token": token})
}

// Register controller handling the register request
func (ctrl *UserController) Register(c *gin.Context) {
	data, err := utils.ParseRequestBody(c, &models.User{})

	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": strings.Split(err.Error(), "\n")})
		return
	}

	user := data.(*models.User)
	if err := userModel.Register(user); err != nil {
		c.AbortWithStatusJSON(http.StatusNotAcceptable, gin.H{"message": strings.Split(err.Error(), "\n")})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully registered", "user": user})
	return
}

// Logout ...
func (ctrl *UserController) Logout(c *gin.Context) {

	au, err := authModel.ExtractTokenMetadata(c.Request)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": "User not logged in"})
		return
	}

	delErr := authModel.DeleteAuth(au.AccessUUID)
	if delErr != nil { // if any goes wrong
		c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"message": "Invalid request"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Successfully logged out"})
}
