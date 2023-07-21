package middleware

import (
	"CTFe/internal/global/redis"
	"CTFe/internal/utils"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GlobalMiddleWare(ctx *gin.Context) {
	// 检测cookies状态
	ctfeToken, err := ctx.Cookie("ctfe_token")
	if err != nil {
		redis.Remove(ctfeToken)
		ctfeToken = utils.GetUUID()
		redis.SetCTFeToken(ctfeToken, 0)
		ctx.SetCookie("ctfe_token", ctfeToken, 36000, "/", "", false, true)
	}

	// 检测白名单
	path := ctx.Request.URL.Path
	if path == "/users/UserRegister" {
		ctx.Next()
	}

	ctfeTokenStatus := redis.GetCTFeToken(ctfeToken)

	// 未授权，重定向至登录页面
	if ctfeTokenStatus != 1 {
		ctx.Redirect(http.StatusMovedPermanently, "https://www.baidu.com")
	}
}
