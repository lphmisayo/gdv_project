package router

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
)

func InitRouter() *gin.Engine {
	gin.SetMode(global.Config.System.Env) //初始化路由环境
	router := gin.Default()               //初始化路由
	apiRouterGroup := router.Group("api")
	SettingsRouter(apiRouterGroup) //系统配置
	ImagesRouter(apiRouterGroup)   //图片接口
	return router
}
