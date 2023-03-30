package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvd_project/models/res"
)

func (s SettingsApi) SettingsInfoView(c *gin.Context) {
	//c.JSON(200, gin.H{"msg": "SettingsInfoTest"})
	res.OK(map[string]string{
		"id": "test01",
	}, "测试信号机制", c)
}
