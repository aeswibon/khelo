package route

import (
	"time"

	"github.com/cp-Coder/khelo/api/controller"
	"github.com/cp-Coder/khelo/bootstrap"
	"github.com/cp-Coder/khelo/domain"
	"github.com/cp-Coder/khelo/mongo"
	"github.com/cp-Coder/khelo/repository"
	"github.com/cp-Coder/khelo/usecase"
	"github.com/gin-gonic/gin"
)

// RefreshTokenRouter ...
func RefreshTokenRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.UserRepository(db, domain.CollectionUser)
	rtc := &controller.RefreshTokenController{
		RefreshTokenUsecase: usecase.RefreshTokenUsecase(ur, timeout),
		Env:                 env,
	}
	group.POST("/refresh", rtc.RefreshToken)
}
