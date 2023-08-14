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

	creatorId, _ := ctx.Get("user_id")

	group.GroupId = rand.Int63()
	preGroup, _ := mysql.SelectGroup(group.GroupId)
	if len(preGroup) > 0 {
		group.GroupId = rand.Int63()
	}

	/*
		校验是否重复创建团队
	*/
	//preParticipations, _ := mysql.SelectParticipation(group.CompetitionId, creatorId)
	//if len(preParticipations) > 0 {
	//	log.ErrorLogger.Println(err.Error())
	//	ctx.JSON(http.StatusBadRequest, response.NewResponse(903, gin.H{"msg": "已经创建该比赛的团队，不能重复创建"}))
	//	ctx.Abort()
	//	return
	//}

	participation := database.Participation{
		ParticipationId: rand.Int63(),
		GroupId:         group.GroupId,
		UserId:          creatorId,
		CompetitionId:   group.CompetitionId,
		IsAdmin:         true,
	}

	err = mysql.CreateGroup(group)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "创建团队失败"}))
		ctx.Abort()
		return
	}

	err = mysql.JoinGroup(participation)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "创建团队失败"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("创建团队[%s]成功", group.GroupName))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "创建成功", "data": group}))
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
