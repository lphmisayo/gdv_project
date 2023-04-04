package settings_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/config"
	"gvd_server/core"
	"gvd_server/global"
	"gvd_server/models/res"
)

func (s SettingsApi) SettingsEmailInfoView(c *gin.Context) {
	res.OKWithData(global.Config.Email, c)
}

func (s SettingsApi) SettingsEmailInfoUpdate(c *gin.Context) {
	var info config.Email
	err := c.ShouldBindJSON(&info)
	if err != nil {
		res.FailWithCode(res.ParamError, c)
		return
	}
	fmt.Println("修改前： ", global.Config.Email)
	global.Config.Email = info
	fmt.Println("修改后： ", global.Config.Email)
	err = core.SetYaml()
	if err != nil {
		global.Log.Errorln(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKBase(c)
}
