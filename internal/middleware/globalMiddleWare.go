package middleware

import (
	"CTFe/internal/models"
	"CTFe/internal/server"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GlobalMiddleWare(ctx *gin.Context) {
	var (
		err       error
		ctfeToken string
	)

	// 检查路由白名单
	pathWhiteList := []string{
		"/users/GetCookies",
		"/users/UserRegister",
		"/users/UserLogin",
	}

	for _, path := range pathWhiteList {
		if path == ctx.FullPath() {
			ctx.Next()
			return
		}
	}

	// 检查cookies是否设置
	ctfeToken, err = ctx.Cookie("ctfe_token")
	if err != nil {
		ctx.JSON(http.StatusForbidden, models.NewResponse(-1, "Cookie未设置"))
		ctx.Abort()
		return
	}

	// 检查登录状态
	ctfeToken, _ = ctx.Cookie("ctfe_token")
	ctfeTokenStatus := server.GetCTFeTokenStatus(ctfeToken)
	if ctfeTokenStatus == 0 {
		ctx.JSON(http.StatusForbidden, models.NewResponse(0, "未登录"))
		ctx.Abort()
		return
	}
}
