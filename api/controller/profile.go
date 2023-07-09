package controller

import (
	"net/http"

	"github.com/cp-Coder/khelo/domain"
	"github.com/gin-gonic/gin"
)

// ProfileController ...
type ProfileController struct {
	ProfileUsecase domain.ProfileUsecase
}

// @BasePath /api

// Fetch ...
// @Summary Fetch profile
// @Description Fetch profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Success 200 {object} domain.User
// @Failure 500 {object} domain.ErrorResponse
// @Router /user/me [get]
func (pc *ProfileController) Fetch(c *gin.Context) {
	userID := c.GetString("x-user-id")

	profile, err := pc.ProfileUsecase.GetProfileByID(c, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, profile)
}

// @BasePath /api

// Update ...
// @Summary Update profile
// @Description Update profile
// @Tags User
// @Accept json
// @Produce json
// @Security ApiKeyAuth
// @Param profile body object true "Profile"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /user/me [patch]
func (pc *ProfileController) Update(c *gin.Context) {
	userID := c.GetString("x-user-id")

	var profile *domain.User
	if err := c.ShouldBind(&profile); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := pc.ProfileUsecase.UpdateProfile(c, userID, profile)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Profile updated successfully"})
}
