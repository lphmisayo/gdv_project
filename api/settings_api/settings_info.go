package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvd_project/global"
	"gvd_project/models/res"
)

type SettingUri struct {
	Name string `uri:"name"`
}

func (s SettingsApi) SettingsInfoView(c *gin.Context) {
	var uri SettingUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.ParamError, c)
		return
	}
	switch uri.Name {
	case "site":
		res.OKWithData(global.Config.SiteInfo, c)
	case "email":
		res.OKWithData(global.Config.Email, c)
	case "qq":
		res.OKWithData(global.Config.QQ, c)
	case "qiniu":
		res.OKWithData(global.Config.Qiniu, c)
	case "jwt":
		res.OKWithData(global.Config.Jwt, c)
	default:
		res.FailWithCode(res.UriError, c)
	}

}
