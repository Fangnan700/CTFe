package service

import (
	"CTFe/server/global/mysql"
	"CTFe/server/middleware"
	"CTFe/server/model/database"
	"CTFe/server/model/response"
	"CTFe/server/util/encrypt"
	"CTFe/server/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"github.com/google/uuid"
	"math/rand"
	"net/http"
	"strconv"
	"time"
)

// UserRegisterHandler 用户注册
func UserRegisterHandler(ctx *gin.Context) {
	var user database.User
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	// 生成随机uuid、id、创建时间
	user.UUID = uuid.New().String()
	user.UserId = rand.Int63()
	user.CreateTime = time.Now().Unix()

	// 检查id、邮箱、手机是否冲突
	preUsers, err := mysql.SelectUser(nil)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "内部异常"}))
		ctx.Abort()
		return
	}
	for _, pu := range preUsers {
		if user.UserId == pu.UserId {
			user.UserId = rand.Int63()
		}
		if user.UserEmail == pu.UserEmail {
			ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "邮箱已注册"}))
			ctx.Abort()
			return
		}
		if user.UserPhone == pu.UserPhone {
			ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "手机号已注册"}))
			ctx.Abort()
			return
		}
	}

	err = mysql.InsertUser(user)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("[%d]注册成功", user.UserId))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "注册成功"}))
}

// DeleteUserHandler 用户删除
func DeleteUserHandler(ctx *gin.Context) {
	var user database.User
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	err = mysql.DeleteUser(user)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "内部异常"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("[%d]删除成功", user.UserId))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "删除成功"}))
}

// UserLoginHandler 用户登录
func UserLoginHandler(ctx *gin.Context) {
	var user database.User
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	auth, err := mysql.SelectUser(user.UserEmail)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}
	if len(auth) <= 0 {
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(901, gin.H{"msg": "账号不存在"}))
		ctx.Abort()
		return
	}

	if user.UserPwd == auth[0].UserPwd {
		if err != nil {
			log.ErrorLogger.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
			ctx.Abort()
			return
		}
		token, _ := middleware.GenToken(encrypt.CopyStr(auth[0].UUID) + "$" + strconv.FormatInt(auth[0].UserId, 10))

		log.InfoLogger.Println(fmt.Sprintf("[%d]登录成功", auth[0].UserId))
		ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "登录成功", "et": token, "uuid": auth[0].UUID, "user_id": auth[0].UserId}))
	} else {
		ctx.JSON(http.StatusOK, response.NewResponse(902, gin.H{"msg": "账号或密码错误"}))
	}
}

// UserLogoutHandler 用户注销
func UserLogoutHandler(ctx *gin.Context) {

}

// UpdateUserHandler 用户更新
func UpdateUserHandler(ctx *gin.Context) {
	var user database.User
	err := ctx.ShouldBindBodyWith(&user, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	// 检查id、邮箱、手机是否冲突
	preUsers, err := mysql.SelectUser(nil)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "内部异常"}))
		ctx.Abort()
		return
	}
	for _, pu := range preUsers {
		if user.UserEmail == pu.UserEmail && user.UserId != pu.UserId {
			ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "邮箱已注册"}))
			ctx.Abort()
			return
		}
		if user.UserPhone == pu.UserPhone && user.UserId != pu.UserId {
			ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "手机号已注册"}))
			ctx.Abort()
			return
		}
	}

	err = mysql.UpdateUser(user)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("[%d]修改成功", user.UserId))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "修改成功"}))
}

// GetUserHandler 查询用户
func GetUserHandler(ctx *gin.Context) {
	var (
		users []database.User
		err   error
	)

	keyword := ctx.Query("keyword")
	if keyword == "" {
		users, err = mysql.SelectUser(nil)
	} else {
		users, err = mysql.SelectUser(keyword)
	}
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	// ！！！过滤敏感信息！！！
	for i, _ := range users {
		users[i].UserPwd = ""
		users[i].UUID = ""
	}

	log.InfoLogger.Println(fmt.Sprintf("查询用户成功"))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "获取成功", "data": users}))
}
