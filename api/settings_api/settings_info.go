package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvd_project/global"
	"gvd_project/models/res"
)

func (s SettingsApi) SettingsInfoView(c *gin.Context) {
	res.OKWithData(global.Config.SiteInfo, c)
}
