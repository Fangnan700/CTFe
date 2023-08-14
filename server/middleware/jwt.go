package middleware

import (
	"CTFe/server/global/mysql"
	"CTFe/server/model/database"
	"CTFe/server/model/response"
	"CTFe/server/util/encrypt"
	"CTFe/server/util/log"
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"strconv"
	"strings"
	"time"
)

type MyClaims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

// TokenExpireDuration 定义过期时间
const TokenExpireDuration = time.Hour * 48

// MySecret 定义secret
var MySecret = []byte("ahuiseal")

// WhiteList 定义白名单
var WhiteList = []string{
	"/user_login",
	"/user_register",
}

// AdminList 定义管理员名单
var AdminList = []string{
	"/add_administrator",
	"/delete_administrator",
	"/get_administrator_list",
	"/create_competition",
	"/delete_competition",
}

// GenToken 生成jwt
func GenToken(username string) (string, error) {
	c := MyClaims{username, jwt.StandardClaims{
		ExpiresAt: time.Now().Add(TokenExpireDuration).Unix(),
		Issuer:    "ctfe",
	},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, c)

	return token.SignedString(MySecret)
}

// ParseToken 解析JWT
func ParseToken(tokenString string) (*MyClaims, error) {
	//解析token
	token, err := jwt.ParseWithClaims(tokenString, &MyClaims{}, func(token *jwt.Token) (i interface{}, err error) {
		return MySecret, nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := token.Claims.(*MyClaims); ok && token.Valid {
		return claims, nil
	}
	return nil, errors.New("invalid token")
}

// JWTAuthMiddleware JWT中间件
func JWTAuthMiddleware() func(c *gin.Context) {
	return func(c *gin.Context) {

		for _, path := range WhiteList {
			if c.FullPath() == path {
				c.Next()
				return
			}
		}

		authHeader := c.Request.Header.Get("Authorization")
		if authHeader == "" {
			c.JSON(http.StatusForbidden, response.NewResponse(902, "未授权"))
			c.Abort()
			return
		}
		parts := strings.SplitN(authHeader, " ", 2)

		if !(len(parts) == 2 && parts[0] == "Bearer") {
			c.JSON(http.StatusForbidden, response.NewResponse(902, "未授权"))
			c.Abort()
			return
		}
		mc, err := ParseToken(parts[1])
		if err != nil {
			c.JSON(http.StatusForbidden, response.NewResponse(902, "未授权"))
			c.Abort()
			return
		}

		var mu database.User
		_ = c.ShouldBindBodyWith(&mu, binding.JSON)

		jwtUSERNAME := strings.Replace(mc.Username, "\"", "", -1)
		reqUSERNAME := encrypt.CopyStr(mu.UUID) + "$" + strconv.FormatInt(mu.UserId, 10)

		if mu.UserId != 0 && reqUSERNAME != jwtUSERNAME {
			c.JSON(http.StatusForbidden, response.NewResponse(902, "未授权"))
			c.Abort()
			return
		}

		for _, path := range AdminList {
			if path == c.FullPath() {
				userId := strings.Split(jwtUSERNAME, "$")[1]
				admins, err := mysql.SelectAdministrator(userId)
				if err != nil {
					log.ErrorLogger.Println(err.Error())
				}
				if len(admins) <= 0 {
					c.JSON(http.StatusForbidden, response.NewResponse(902, "未授权"))
					c.Abort()
					return
				}
			}
		}

		c.Set("user_id", strings.Split(jwtUSERNAME, "$")[1])
		c.Next()
	}

}
