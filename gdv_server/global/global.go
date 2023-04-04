package global

import (
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
	"gvd_server/config"
	"gvd_server/models"
)

var (
	Config         *config.Config
	DB             *gorm.DB
	Log            *logrus.Logger
	MysqlLogConfig *models.MysqlLogSt //细化MysqlDB日志操作
)
