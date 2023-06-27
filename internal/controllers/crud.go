package controllers

import (
	"net/http"
	"strconv"
	"strings"

	db "github.com/cp-Coder/khelo/pkg/platform/database"
	"github.com/cp-Coder/khelo/pkg/utils"
	"github.com/gin-gonic/gin"
)

// CRUDController struct to hold the CRUD controller
type CRUDController struct {
	// Model on which the CRUD operations are to be performed
	Model interface{}
}

// Init method to migrate the schema
func (ctrl *CRUDController) Init() {
	db.GetDBClient().Migrate(ctrl.Model)
}

// GetAll method to get all records
func (ctrl *CRUDController) GetAll(c *gin.Context) {
	var records []interface{}

	// Get the query params
	query := c.Request.URL.Query()

	// Get the limit
	limitStr := c.DefaultQuery("limit", "10")
	// Parse limit and offset values
	limit, _ := strconv.Atoi(limitStr)

	// Get the offset
	pageStr := c.DefaultQuery("page", "1")
	page, _ := strconv.Atoi(pageStr)
	offset := (page - 1) * limit

	// set the model
	model := db.GetDBClient().Model(ctrl.Model)

	result := model.Limit(limit).Offset(offset).Find(&records)
	if result.Error != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
		return
	}

	for key, value := range query {
		result = model.Where(key+" = ?", value).Find(&records)
		if result.Error != nil {
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": result.Error.Error()})
			return
		}
	}

	result.Limit(limit).Offset(offset).Find(&records)
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"data": records})
	return
}

// GetByID method to get a record by id
func (ctrl *CRUDController) GetByID(c *gin.Context) {
	var record interface{}

	// set the model
	model := db.GetDBClient().Model(ctrl.Model)
	if err := model.First(&record, c.Param("id")); err != nil {
		c.JSON(500, gin.H{"error": err})
		return
	}
	c.JSON(200, gin.H{"record": record})
}

// Create method to create a record
func (ctrl *CRUDController) Create(c *gin.Context) {
	data, err := utils.ParseRequestBody(c, ctrl.Model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": strings.Split(err.Error(), "\n")})
		return
	}
	if err := db.GetDBClient().Create(data); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusCreated, gin.H{"data": data})
}

// Update method to update a record
func (ctrl *CRUDController) Update(c *gin.Context) {
	data, err := utils.ParseRequestBody(c, ctrl.Model)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, gin.H{"message": strings.Split(err.Error(), "\n")})
		return
	}

	if err := db.GetDBClient().Update(data, c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusOK, gin.H{"data": data})
	return
}

// Delete method to delete a record
func (ctrl *CRUDController) Delete(c *gin.Context) {
	var record interface{}
	if err := db.GetDBClient().Delete(record, c.Param("id")); err != nil {
		c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.AbortWithStatusJSON(http.StatusNoContent, gin.H{"data": record})
	return
}

// Routes method to register the routes
func Routes(r *gin.RouterGroup, path string, ctrl *CRUDController) {
	// Get all records
	r.GET(path+"/", ctrl.GetAll)

	// Get a record by id
	r.GET(path+"/:id", ctrl.GetByID)

	// Create a record
	r.POST(path+"/", ctrl.Create)

	// Update a record
	r.PUT(path+"/:id", ctrl.Update)

	// Delete a record
	r.DELETE(path+"/:id", ctrl.Delete)
}
