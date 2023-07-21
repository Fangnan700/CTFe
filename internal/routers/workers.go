package routers

import (
	"CTFe/internal/models"
	"CTFe/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

// UserRegister 用户注册
func UserRegister(ctx *gin.Context) {
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
	}

	resp.Code = 1
	resp.Body = users

	ctx.JSON(http.StatusOK, resp)
}
