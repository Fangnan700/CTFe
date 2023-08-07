package routers

import (
	"CTFe/internal/middleware"
	"github.com/gin-gonic/gin"
)

var (
	engine      *gin.Engine
	usersRouter *gin.RouterGroup
)

func Register() error {
	engine = gin.Default()

	// 注册全局中间件
	engine.Use(middleware.GlobalMiddleWare)

	// 用户相关的路由
	usersRouter = engine.Group("/users")
	{
		usersRouter.GET("/GetCookies", GetCookies)
		usersRouter.GET("/QueryUsers", QueryUsers)
		usersRouter.GET("/QueryAllUsers", QueryAllUsers)
		usersRouter.POST("/UserRegister", UserRegister)
		usersRouter.POST("/UserLogin", UserLogin)
		usersRouter.POST("/UserLogout", UserLogout)
		usersRouter.PUT("UpdateUser", UpdateUser)
		usersRouter.DELETE("/DeleteUser", DeleteUser)
	}

	err := engine.Run("0.0.0.0:8080")
	if err != nil {
		return err
	}
	return nil
}
