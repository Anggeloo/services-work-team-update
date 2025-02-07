package routes

import (
	"services-work-team-update/controllers"

	"github.com/gin-gonic/gin"
)

func WorkTeamRoutes(r *gin.RouterGroup) {
	r.PUT("/work-team/update/:teamCode", controllers.UpdateWorkTeam)
}
