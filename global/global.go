package global

import (
	"gorm.io/gorm"
	"gvd_project/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
)
