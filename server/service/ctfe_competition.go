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

// CreateCompetitionHandler 创建比赛
func CreateCompetitionHandler(ctx *gin.Context) {
	var competition database.Competition
	err := ctx.ShouldBindBodyWith(&competition, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	competition.CompetitionId = rand.Int63()
	preCompetitions, _ := mysql.SelectCompetition(competition.CompetitionId)
	if len(preCompetitions) > 0 {
		competition.CompetitionId = rand.Int63()
	}

	err = mysql.InsertCompetition(competition)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "创建失败"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("创建比赛[%s]成功", competition.CompetitionName))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "创建成功", "data": competition}))
}

// DeleteCompetitionHandler 删除比赛
func DeleteCompetitionHandler(ctx *gin.Context) {
	var competition database.Competition
	err := ctx.ShouldBindBodyWith(&competition, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	competitions, err := mysql.SelectCompetition(competition.CompetitionId)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "系统异常"}))
		ctx.Abort()
		return
	}
	competition = competitions[0]

	err = mysql.DeleteCompetition(competition)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "删除失败"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("删除比赛[%s]成功", competition.CompetitionName))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "删除成功"}))
}

// UpdateCompetitionHandler 更新比赛
func UpdateCompetitionHandler(ctx *gin.Context) {
	var competition database.Competition
	err := ctx.ShouldBindBodyWith(&competition, binding.JSON)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "参数异常"}))
		ctx.Abort()
		return
	}

	err = mysql.UpdateCompetition(competition)
	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "更新失败"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println(fmt.Sprintf("更新比赛[%s]成功", competition.CompetitionName))
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "更新成功", "data": competition}))
}

// SelectCompetitionHandler 查询比赛
func SelectCompetitionHandler(ctx *gin.Context) {
	var competitions []database.Competition
	var err error
	keyword := ctx.Query("keyword")
	if keyword == "" {
		competitions, err = mysql.SelectCompetition(nil)
	} else {
		competitions, err = mysql.SelectCompetition(keyword)
	}

	if err != nil {
		log.ErrorLogger.Println(err.Error())
		ctx.JSON(http.StatusInternalServerError, response.NewResponse(903, gin.H{"msg": "查询失败"}))
		ctx.Abort()
		return
	}

	log.InfoLogger.Println("查询比赛成功")
	ctx.JSON(http.StatusOK, response.NewResponse(900, gin.H{"msg": "查询成功", "data": competitions}))
}
