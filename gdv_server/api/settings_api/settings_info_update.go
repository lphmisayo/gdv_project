package settings_api

import (
	"github.com/gin-gonic/gin"
	"gvd_project/config"
	"gvd_project/core"
	"gvd_project/global"
	"gvd_project/models/res"
)

// SettingsInfoUpdate 多
func (s SettingsApi) SettingsInfoUpdate(c *gin.Context) {
	var uri SettingUri
	err := c.ShouldBindUri(&uri)
	if err != nil {
		res.FailWithCode(res.ParamError, c)
		return
	}
	switch uri.Name {
	case "site":
		var info config.SiteInfo
		err = c.ShouldBindJSON(&info)
		if HandleErr(err, res.ParamError, c) {
			return
		}
		global.Config.SiteInfo = info
	case "email":
		var info config.Email
		err = c.ShouldBindJSON(&info)
		if HandleErr(err, res.ParamError, c) {
			return
		}
		global.Config.Email = info
	case "qq":
		var info config.QQ
		err = c.ShouldBindJSON(&info)
		if HandleErr(err, res.ParamError, c) {
			return
		}
		global.Config.QQ = info
	case "jwt":
		var info config.Jwt
		err = c.ShouldBindJSON(&info)
		if HandleErr(err, res.ParamError, c) {
			return
		}
		global.Config.Jwt = info
	case "qiniu":
		var info config.Qiniu
		err = c.ShouldBindJSON(&info)
		if HandleErr(err, res.ParamError, c) {
			return
		}
		global.Config.Qiniu = info
	}

	//fmt.Println("修改前： ", global.Config.SiteInfo)

	//fmt.Println("修改后： ", global.Config.SiteInfo)
	UpdateYaml(c)
}

func HandleErr(err error, errType res.ErrCode, c *gin.Context) bool {
	if err != nil {
		res.FailWithCode(errType, c)
		return true
	}
	return false
}

func UpdateYaml(c *gin.Context) {
	err := core.SetYaml()
	if err != nil {
		global.Log.Errorln(err.Error())
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKBase(c)
}
