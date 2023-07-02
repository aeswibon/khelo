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

// AuthRouter router defining all auth routes
func AuthRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	ur := repository.UserRepository(db, domain.CollectionUser)
	sc := controller.RegisterController{
		RegisterUsecase: usecase.RegisterUsecase(ur, timeout),
		Env:             env,
	}
	lc := controller.LoginController{
		LoginUsecase: usecase.LoginUsecase(ur, timeout),
		Env:          env,
	}

	group.POST("/register", sc.Register)
	group.POST("/login", lc.Login)

}
