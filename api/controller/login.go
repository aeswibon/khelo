package controller

import (
	"net/http"

	"github.com/cp-Coder/khelo/bootstrap"
	"github.com/cp-Coder/khelo/domain"
	"github.com/gin-gonic/gin"
)

// LoginController ...
type LoginController struct {
	LoginUsecase domain.LoginUsecase
	Env          *bootstrap.Env
}

// @BasePath /api

// Login method defining login
// @Summary Login a user
// @Description Login a user
// @Tags auth
// @Accept json
// @Produce json
// @Param login body domain.LoginRequest true "Login"
// @Success 200 {object} domain.LoginResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /auth/login [post]
func (lc *LoginController) Login(c *gin.Context) {
	var request domain.LoginRequest

	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	user, err := lc.LoginUsecase.Authenticate(c, request)
	if err != nil {
		c.JSON(http.StatusNotFound, domain.ErrorResponse{Message: err.Error()})
		return
	}

	accessToken, err := lc.LoginUsecase.CreateAccessToken(&user, lc.Env.AccessTokenSecret, lc.Env.AccessTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	refreshToken, err := lc.LoginUsecase.CreateRefreshToken(&user, lc.Env.RefreshTokenSecret, lc.Env.RefreshTokenExpiryHour)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	loginResponse := domain.LoginResponse{
		AccessToken:  accessToken,
		RefreshToken: refreshToken,
	}

	c.JSON(http.StatusOK, loginResponse)
	return
}
