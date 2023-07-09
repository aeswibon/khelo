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

// FacilityRouter ...
func FacilityRouter(env *bootstrap.Env, timeout time.Duration, db mongo.Database, group *gin.RouterGroup) {
	fr := repository.FacilityRepository(db, domain.CollectionFacility)
	fc := &controller.FacilityController{
		FacilityUsecase: usecase.FacilityUsecase(fr, timeout),
	}
	group.GET("/facility", fc.Fetch)
	group.POST("/facility", fc.Create)
	group.GET("/facility/:id", fc.GetFacilityByID)
	group.GET("/facility/name/:name", fc.GetFacilityByName)
	group.PATCH("/facility/:id", fc.Update)
}
