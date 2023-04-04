package api

import (
	"gvd_server/api/images_api"
	"gvd_server/api/settings_api"
)

type ApiGroup struct {
	SettingsApi settings_api.SettingsApi
	ImageApi    images_api.ImageApi
}

var ApiGroupApp = new(ApiGroup)
