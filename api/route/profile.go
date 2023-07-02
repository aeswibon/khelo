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

// ProfileRouter ...
func ProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.UserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.ProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
	group.PATCH("/profile", pc.Update)
}
