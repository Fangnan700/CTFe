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
	用户注册、登录、登出、注销模块
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
		ctx.JSON(http.StatusBadRequest, models.NewResponse(-1, "参数异常"))
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
		ctx.JSON(http.StatusBadRequest, models.NewResponse(-1, "参数异常"))
		return
	}

	loginErr = server.UserLoginServer(loginInfo)
	if loginErr != nil {
		ctx.JSON(http.StatusOK, models.NewResponse(-1, loginErr.Message))
	}

	ctfeToken, _ := ctx.Cookie("ctfe_token")
	server.SetCTFeTokenStatus(ctfeToken, 1)

	ctx.JSON(http.StatusOK, models.NewResponse(1, "登录成功"))
	return
}

// UserLogout 用户登出
func UserLogout(ctx *gin.Context) {
	ctfeToken, _ := ctx.Cookie("ctfe_token")
	server.RemoveCTFeTokenStatus(ctfeToken)
	ctx.JSON(http.StatusOK, models.NewResponse(1, "登出成功"))
}

// DeleteUser 用户注销
func DeleteUser(ctx *gin.Context) {
	var user models.Users
	_ = ctx.ShouldBindJSON(&user)

	err := server.DeleteUserServer(user.UserId, user.UserPwd)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, err)
		return
	}

	ctx.JSON(http.StatusOK, models.NewResponse(1, "注销成功"))
}

/*
	用户数据查询模块
*/

// QueryAllUsers 查询所有用户
func QueryAllUsers(ctx *gin.Context) {
	users, ctfeError := server.QueryAllUsersServer()
	if ctfeError != nil {
		ctx.JSON(http.StatusInternalServerError, "查询异常")
		return
	}

	ctx.JSON(http.StatusOK, users)
}

func QueryUsers(ctx *gin.Context) {
	keyword := ctx.Query("keyword")
	users, ctfeError := server.QueryUsers(keyword)
	if ctfeError != nil {
		ctx.JSON(http.StatusInternalServerError, "查询异常")
		return
	}

	ctx.JSON(http.StatusOK, users)
}

/*
	用户信息更新模块
*/

// UpdateUser 更新用户信息
func UpdateUser(ctx *gin.Context) {
	var user models.Users
	_ = ctx.ShouldBindJSON(&user)

	err := server.UpdateUserServer(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.NewResponse(-1, "更新异常："+err.Message))
		return
	}

	ctx.JSON(http.StatusBadRequest, models.NewResponse(1, "更新成功"))
	return
}
