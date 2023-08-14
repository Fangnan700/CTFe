package v1

import (
	"CTFe/server/service"
	"github.com/gin-gonic/gin"
)

func CTFeUserRegisterRoute(r *gin.Engine) {
	r.POST("/user_login", service.UserLoginHandler)
	r.POST("/user_logout", service.UserLogoutHandler)
	r.POST("/user_register", service.UserRegisterHandler)
	r.DELETE("/delete_user", service.DeleteUserHandler)
	r.PUT("/update_user", service.UpdateUserHandler)
	r.GET("/get_user", service.GetUserHandler)
}
