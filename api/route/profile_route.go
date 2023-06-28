package route

import (
	"time"

	"github.com/gin-gonic/gin"
	"gitlab.com/cp-Coder/khelo/api/controller"
	"gitlab.com/cp-Coder/khelo/bootstrap"
	"gitlab.com/cp-Coder/khelo/domain"
	"gitlab.com/cp-Coder/khelo/mongo"
	"gitlab.com/cp-Coder/khelo/repository"
	"gitlab.com/cp-Coder/khelo/usecase"
)

func NewProfileRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.NewUserRepository(db, domain.CollectionUser)
	pc := &controller.ProfileController{
		ProfileUsecase: usecase.NewProfileUsecase(ur, timeout),
	}
	group.GET("/profile", pc.Fetch)
}
