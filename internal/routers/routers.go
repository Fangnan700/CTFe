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
		// 查询所有用户
		usersRouter.GET("/QueryAllUsers", QueryAllUsers)

		// 用户注册
		usersRouter.POST("/UserRegister", UserRegister)

		// 用户登录
		usersRouter.POST("/UserLogin", UserLogin)
	}

	err := engine.Run("0.0.0.0:8080")
	if err != nil {
		return err
	}
	return nil
}
