package router

import (
	"github.com/gin-gonic/gin"
	"gvd_server/api"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingApi := api.ApiGroupApp.SettingsApi
	{
		router.GET("settings/:name", settingApi.SettingsInfoView)
		router.PUT("settings/:name", settingApi.SettingsInfoUpdate)
	}
}
