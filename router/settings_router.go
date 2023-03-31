package router

import (
	"github.com/gin-gonic/gin"
	"gvd_project/api"
)

func SettingsRouter(router *gin.RouterGroup) {
	settingApi := api.ApiGroupApp.SettingsApi
	{
		router.GET("settings", settingApi.SettingsInfoView)
		router.PUT("settings", settingApi.SettingsInfoUpdate)
	}
}
