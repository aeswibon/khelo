package middleware

import (
	"github.com/cp-Coder/khelo/internal/controllers"
	"github.com/gin-gonic/gin"
)

// TokenAuthMiddleware ...
// JWT Authentication middleware attached to each request that needs to be authenitcated to validate the access_token in the header
func TokenAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		auth := &controllers.AuthController{}
		auth.TokenValid(c)
		c.Next()
	}
}
