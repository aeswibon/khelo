package controller

import (
	"net/http"

	"github.com/cp-Coder/khelo/bootstrap"
	"github.com/cp-Coder/khelo/domain"
	"github.com/gin-gonic/gin"
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

	if err := sc.RegisterUsecase.Create(c, &request); err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	RegisterResponse := domain.RegisterResponse{
		Message: "User created successfully",
	}

	c.JSON(http.StatusOK, RegisterResponse)
}
