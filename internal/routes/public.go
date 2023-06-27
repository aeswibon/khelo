package routes

import (
	"github.com/cp-Coder/khelo/internal/controllers"
	"github.com/cp-Coder/khelo/internal/models"
	"github.com/gin-gonic/gin"
)

// PublicRoutes ...
func PublicRoutes(r *gin.RouterGroup) {
	user := &controllers.UserController{}

	state := &controllers.CRUDController{Model: &models.State{}}
	state.Init()

	district := &controllers.
		CRUDController{Model: &models.District{}}
	district.Init()

	r.POST("/user/login", user.Login)
	r.POST("/user/register", user.Register)
	r.GET("/user/logout", user.Logout)

	// State routes
	controllers.Routes(r, "state", state)

	// District routes
	controllers.Routes(r, "district", district)
}
