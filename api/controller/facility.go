package controller

import (
	"net/http"

	"github.com/cp-Coder/khelo/domain"
	"github.com/gin-gonic/gin"
)

// FacilityController ...
type FacilityController struct {
	FacilityUsecase domain.FacilityUsecase
}

// @BasePath /api

// Fetch ...
// @Summary Fetch all facilities
// @Description Fetch all facilities
// @Tags Facility
// @Accept json
// @Produce json
// @Success 200 {array} domain.Facility
// @Failure 500 {object} domain.ErrorResponse
// @Router /facility [get]
func (fc *FacilityController) Fetch(c *gin.Context) {
	facilities, err := fc.FacilityUsecase.Fetch(c)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, facilities)
}

// @BasePath /api

// Create ...
// @Summary Create a facility
// @Description Create a facility
// @Tags Facility
// @Accept json
// @Produce json
// @Param facility body domain.Facility true "Facility"
// @Success 201 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /facility [post]
func (fc *FacilityController) Create(c *gin.Context) {
	var facility *domain.Facility
	if err := c.ShouldBind(&facility); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := fc.FacilityUsecase.Create(c, facility)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusCreated, domain.SuccessResponse{Message: "Facility created successfully"})
}

// @BasePath /api

// GetFacilityByName ...
// @Summary Get a facility by name
// @Description Get a facility by name
// @Tags Facility
// @Accept json
// @Produce json
// @Param name path string true "Facility Name"
// @Success 200 {object} domain.Facility
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /facility/{name} [get]
func (fc *FacilityController) GetFacilityByName(c *gin.Context) {
	name := c.Param("name")

	facility, err := fc.FacilityUsecase.GetFacilityByName(c, name)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, facility)
}

// @BasePath /api

// GetFacilityByID ...
// @Summary Get a facility by id
// @Description Get a facility by id
// @Tags Facility
// @Accept json
// @Produce json
// @Param id path string true "Facility ID"
// @Success 200 {object} domain.Facility
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /facility/{id} [get]
func (fc *FacilityController) GetFacilityByID(c *gin.Context) {
	id := c.Param("id")

	facility, err := fc.FacilityUsecase.GetFacilityByID(c, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, facility)
}

// @BasePath /api

// Update ...
// @Summary Update a facility
// @Description Update a facility
// @Tags Facility
// @Accept json
// @Produce json
// @Param id path string true "Facility ID"
// @Param facility body domain.Facility true "Facility"
// @Success 200 {object} domain.SuccessResponse
// @Failure 400 {object} domain.ErrorResponse
// @Failure 500 {object} domain.ErrorResponse
// @Router /facility/{id} [patch]
func (fc *FacilityController) Update(c *gin.Context) {
	id := c.Param("id")

	var facility *domain.Facility
	if err := c.ShouldBind(&facility); err != nil {
		c.JSON(http.StatusBadRequest, domain.ErrorResponse{Message: err.Error()})
		return
	}

	err := fc.FacilityUsecase.UpdateFacility(c, id, facility)
	if err != nil {
		c.JSON(http.StatusInternalServerError, domain.ErrorResponse{Message: err.Error()})
		return
	}

	c.JSON(http.StatusOK, domain.SuccessResponse{Message: "Facility updated successfully"})
}
