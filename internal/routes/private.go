package routes

import (
	"github.com/cp-Coder/khelo/internal/controllers"
	"github.com/gin-gonic/gin"
)

// ServiceLocator Specifies the dependencies for each controller
type ServiceLocator struct {
	AuthController *controllers.AuthController
	// Add other controller instances here
}

// PrivateRoutes ...
func PrivateRoutes(r *gin.RouterGroup, locator *ServiceLocator) {
	auth := locator.AuthController

	//Refresh the token when needed to generate new access_token and refresh_token for the user
	r.POST("/token/refresh", auth.Refresh)

	// // Grouped routes that require authentication
	// authenticated := r.Group("/", middleware.TokenAuthMiddleware())
	// {
	// }

}
