package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/cp-Coder/khelo/bootstrap"
	"gitlab.com/cp-Coder/khelo/mongo"
)

// AuthRouter router defining all auth routes
func AuthRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {

}
