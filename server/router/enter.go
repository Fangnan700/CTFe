package router

import (
	"CTFe/server/global/config"
	"CTFe/server/middleware"
	v1 "CTFe/server/router/v1"
	"CTFe/server/util/log"
	"fmt"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

func init() {
	gin.SetMode(gin.ReleaseMode)

	router = gin.New()
	router.Use(middleware.JWTAuthMiddleware())

	v1.CTFeAdminRegisterRoute(router)
	v1.CTFeChallengeRegisterRoute(router)
	v1.CTFeCompetitionRegisterRoute(router)
	v1.CTFeContainerRegisterRoute(router)
	v1.CTFeGroupRegisterRoute(router)
	v1.CTFeParticipationRegisterRoute(router)
	v1.CTFeScoreRegisterRoute(router)
	v1.CTFeSolvedRegisterRoute(router)
	v1.CTFeUserRegisterRoute(router)
}

func Start() {
	err := router.Run(fmt.Sprintf("%s:%d", config.GlobalConfig.ServerConfig.Host, config.GlobalConfig.ServerConfig.Port))
	if err != nil {
		log.ErrorLogger.Println(err)
	}
}
