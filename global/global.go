package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvd_project/config"
)

var (
	Config *config.Config
	DB     *gorm.DB
	Log    *logrus.Logger
)
