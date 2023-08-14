package service

import (
	"CTFe/server/global/mysql"
	"CTFe/server/model/database"
	"CTFe/server/model/response"
	"CTFe/server/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"math/rand"
	"net/http"
)

// AddAdministratorHandler 添加管理员
func AddAdministratorHandler(ctx *gin.Context) {
	var jsonData map[string]int64

	err := ctx.ShouldBindBodyWith(&jsonData, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	var admin database.Admin
	admin.AdminId = rand.Int63()
	admin.UserId = jsonData["target_user_id"]
	err = mysql.InsertAdministrator(admin)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "系统异常"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("添加管理员[%d]成功", admin.UserId))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "添加成功"}))
}

// DeleteAdministratorHandler 删除管理员
func DeleteAdministratorHandler(ctx *gin.Context) {
	var jsonData map[string]int64

	err := ctx.ShouldBindBodyWith(&jsonData, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	var admin database.Admin
	admin.AdminId = jsonData["target_admin_id"]
	admin.UserId = jsonData["target_user_id"]
	err = mysql.DeleteAdministrator(admin)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "系统异常"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("删除管理员[%d]成功", admin.UserId))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "删除成功"}))
}

// SelectAdministratorHandler 查询管理员
func SelectAdministratorHandler(ctx *gin.Context) {
	var (
		admins []database.Admin
		err    error
	)

	id := ctx.Query("id")
	if id == "" {
		admins, err = mysql.SelectAdministrator(nil)
	} else {
		admins, err = mysql.SelectAdministrator(id)
	}
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	var results = make(map[int64]database.User)
	for i, _ := range admins {
		users, err := mysql.SelectUser(admins[i].UserId)
		if err != nil {
			log.ErrorLogger.Println(err.Error())
			ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
			ctx.Abort()
			return
		}
		for _, u := range users {
			u.UUID = ""
			u.UserPwd = ""
			results[admins[i].AdminId] = u
		}
	}

	log.InfoLogger.Println(fmt.Sprintf("查询管理员成功"))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "获取成功", "data": results}))
}
