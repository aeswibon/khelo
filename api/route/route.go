package route

import (
	"time"

	"github.com/cp-Coder/khelo/api/middleware"
	"github.com/cp-Coder/khelo/bootstrap"
	"github.com/cp-Coder/khelo/mongo"
	"github.com/gin-gonic/gin"
)

// @title Khelo API
// @version 1.0
// @description This is a sample server celler server.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @BasePath /api

// Setup function to setup all routes
func Setup(env *bootstrap.Env, timeout time.Duration, db mongo.Database, gin *gin.Engine) {
	publicRouter := gin.Group("/api")
	// All Public APIs
	AuthRouter(env, timeout, db, publicRouter)
	RefreshTokenRouter(env, timeout, db, publicRouter)

	protectedRouter := gin.Group("")
	// Middleware to verify AccessToken
	protectedRouter.Use(middleware.JwtAuthMiddleware(env.AccessTokenSecret))
	// All Private APIs
	ProfileRouter(env, timeout, db, protectedRouter)
	FacilityRouter(env, timeout, db, protectedRouter)
}
