package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_project/config"
	"gvd_project/core"
	"gvd_project/global"
	"gvd_project/models/res"
)

func (s SettingsApi) SettingsInfoUpdate(c *gin.Context) {
	var info config.SiteInfo
	err := c.ShouldBindJSON(&info)
	if err != nil {
		res.FailWithCode(res.ParamError, c)
		return
	}
	fmt.Println("修改前： ", global.Config.SiteInfo)
	global.Config.SiteInfo = info
	fmt.Println("修改后： ", global.Config.SiteInfo)
	err = core.SetYaml()
	if err != nil {
		global.Log.Errorln(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKBase(c)
}
