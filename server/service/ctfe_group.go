package service

import (
	"CTFe/server/model/database"
	"CTFe/server/model/response"
	"CTFe/server/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
)

// CreateGroupHandler 创建团队
func CreateGroupHandler(ctx *gin.Context) {
	var group database.Group

	err := ctx.ShouldBindBodyWith(&group, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	fmt.Println(ctx.Get("user_id"))
	fmt.Printf("%+v\n", group)

	ctx.JSON(http.StatusOK, response.NewResponse(903, gin.H{"msg": "创建成功", "data": group}))
}

// DeleteGroupHandler 删除团队
func DeleteGroupHandler(ctx *gin.Context) {

}

// UpdateGroupHandler 更新团队
func UpdateGroupHandler(ctx *gin.Context) {

}

// SelectGroupHandler 查询团队
func SelectGroupHandler() {

}
