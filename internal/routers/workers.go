package routers

import (
	"CTFe/internal/models"
	"CTFe/internal/server"
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
	var err error
	var resp models.Response
	var registerInfo models.Users

	err = ctx.ShouldBindJSON(&registerInfo)
	if err != nil {
		resp.Code = -1
		resp.Body = "参数异常"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	err = server.UserRegisterServer(registerInfo)
	if err != nil {
		fmt.Println(err)
		resp.Code = -1
		resp.Body = "注册异常"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	// 注册成功，重定向至登录页
	ctx.Redirect(http.StatusMovedPermanently, "")
	return
}

// UserLogin 用户登录
func UserLogin(ctx *gin.Context) {
	var resp models.Response
	
	ctx.JSON(http.StatusOK, resp)
}

// QueryAllUsers 查询所有用户
func QueryAllUsers(ctx *gin.Context) {
	var resp models.Response

	users, err := server.SelectAllUsers()
	if err != nil {
		resp.Code = -1
		resp.Body = "查询异常"
		ctx.JSON(http.StatusOK, resp)
		return
	}

	resp.Code = 1
	resp.Body = users
	ctx.JSON(http.StatusOK, resp)
	return
}
