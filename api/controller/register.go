package controller

import (
	"net/http"

	"github.com/cp-Coder/khelo/bootstrap"
	"github.com/cp-Coder/khelo/domain"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

// RegisterController ...
type RegisterController struct {
	RegisterUsecase domain.RegisterUsecase
	Env             *bootstrap.Env
}

// @BasePath /api

// Register method defining registration
// @Summary Register a new user
// @Description Register a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body domain.RegisterRequest true "Register"
// @Success 200 {object} domain.RegisterResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /auth/register [post]
func (sc *RegisterController) Register(c *gin.Context) {
	var request domain.RegisterRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	if _, err := sc.RegisterUsecase.CheckUser(c, request.Username, request.Email); err == nil {
		c.JSON(http.StatusConflict, domain.ErrorResponse{Message: "User already exists !!!"})
		return
	}

	encryptedPassword, err := bcrypt.GenerateFromPassword(
		[]byte(request.Password),
		bcrypt.DefaultCost,
	)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	request.Password = string(encryptedPassword)

	user := domain.User{
		ID:       primitive.NewObjectID(),
		Username: request.Username,
		Email:    request.Email,
		Password: request.Password,
		Type:     "USER",
	}

	if err := sc.RegisterUsecase.Create(c, &user); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	RegisterResponse := domain.RegisterResponse{
		Message: "User created successfully",
	}

	c.JSON(http.StatusOK, RegisterResponse)
}
