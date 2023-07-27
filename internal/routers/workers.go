package routers

import (
	"CTFe/internal/models"
	"CTFe/internal/server"
	"CTFe/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

// GetCookies 设置cookie
func GetCookies(ctx *gin.Context) {
	ctfeToken := utils.GetUUID()
	ctx.SetCookie("ctfe_token", ctfeToken, 36000, "/", "", false, false)
	ctx.JSON(http.StatusOK, models.NewResponse(1, "获取Cookies成功"))
}

/*
	用户注册、登录模块
*/

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var (
		err          error
		registerErr  *models.CTFeError
		registerInfo models.Users
	)

	err = ctx.ShouldBindJSON(&registerInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResponse(-1, "参数异常"))
		return
	}

	registerErr = server.UserRegisterServer(registerInfo)
	if registerErr != nil {
		ctx.JSON(http.StatusOK, models.NewResponse(-1, registerErr.Message))
		return
	}

	// 注册成功
	ctx.JSON(http.StatusOK, models.NewResponse(1, "注册成功"))
	return
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var (
		err       error
		loginErr  *models.CTFeError
		loginInfo models.Users
	)

	err = ctx.ShouldBindJSON(&loginInfo)
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResponse(-1, "参数异常"))
		return
	}

	loginErr = server.UserLoginServer(loginInfo)
	if loginErr != nil {
		ctx.JSON(http.StatusOK, models.NewResponse(-1, loginErr.Message))
	}

	// 登录成功
	ctfeToken, _ := ctx.Cookie("ctfe_token")
	server.SetCTFeTokenStatus(ctfeToken, 1)
	ctx.JSON(http.StatusOK, models.NewResponse(1, "登录成功"))
	return
}

/*
	用户数据查询模块
*/

// QueryAllUsers 查询所有用户
func QueryAllUsers(ctx *gin.Context) {
	users, err := server.SelectAllUsers()
	if err != nil {
		ctx.JSON(http.StatusOK, models.NewResponse(-1, "查询异常"))
		return
	}
	ctx.JSON(http.StatusOK, models.NewResponse(1, users))
	return
}
