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

func NewTaskRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	tr := repository.NewTaskRepository(db, domain.CollectionTask)
	tc := &controller.TaskController{
		TaskUsecase: usecase.NewTaskUsecase(tr, timeout),
	}
	group.GET("/task", tc.Fetch)
	group.POST("/task", tc.Create)
}
