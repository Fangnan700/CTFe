package v1

import (
	"CTFe/server/service"
	"github.com/gin-gonic/gin"
)

func CTFeCompetitionRegisterRoute(r *gin.Engine) {
	r.POST("/create_competition", service.CreateCompetitionHandler)
	r.DELETE("/delete_competition", service.DeleteCompetitionHandler)
	r.PUT("/update_competition", service.UpdateCompetitionHandler)
	r.GET("/get_competition", service.SelectCompetitionHandler)
}
