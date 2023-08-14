package v1

import (
	"CTFe/server/service"
	"github.com/gin-gonic/gin"
)

func CTFeAdminRegisterRoute(r *gin.Engine) {
	r.GET("/get_administrator_list", service.SelectAdministratorHandler)
	r.POST("/add_administrator", service.AddAdministratorHandler)
	r.DELETE("/delete_administrator", service.DeleteAdministratorHandler)
}
