package v1

import (
	"CTFe/server/service"
	"github.com/gin-gonic/gin"
)

func CTFeGroupRegisterRoute(r *gin.Engine) {
	r.POST("create_group", service.CreateGroupHandler)
}
