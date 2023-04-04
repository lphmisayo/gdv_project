package api

import (
	"gvd_project/api/images_api"
	"gvd_project/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    images_api.ImageApi
}

var ApiGroupApp = new(ApiGroup)
