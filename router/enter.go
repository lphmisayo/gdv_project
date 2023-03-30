package router

import (
	"github.com/gin-gonic/gin"
	"gvd_project/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env)
	router := gin.Default()
	apiRouterGroup := router.Group("api")
	SettingsRouter(apiRouterGroup) //系统配置
	return router
}
