package core

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gvd_server/global"
	"gvd_server/models"
	"time"
)

// &parseTime=Truc&loc=Local
func InitGorm() *gorm.DB {
	if global.Config.Mysql.Host == "" {
		//global.Log.Warn("未配置mysql， 取消gorm连接")
		global.Log.Warnln("未配置mysql， 取消gorm连接")
		return nil
	}
	dsn := global.Config.Mysql.Dsn()

	var mysqlLogger logger.Interface
	var mysqlLogSt models.MysqlLogSt
	if global.Config.System.Env == "debug" {
		mysqlLogger = logger.Default.LogMode(logger.Info)
		mysqlLogSt.Debug = true
	} else {
		mysqlLogger = logger.Default.LogMode(logger.Error) // 只打印错误的sql
		mysqlLogSt.Debug = false
	}
	mysqlLogSt.MysqlLog = logger.Default.LogMode(logger.Info)
	global.MysqlLogConfig = &mysqlLogSt

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: mysqlLogger,
	})
	if err != nil {
		//global.Log.Error(fmt.Sprintf("[%s] mysql连接失败", dsn))
		global.Log.Errorln(fmt.Sprintf("[%s] mysql连接失败", dsn))
		panic(err)
	}

	sqlDB, _ := db.DB()
	sqlDB.SetMaxIdleConns(10)               //最大空闲连接数
	sqlDB.SetMaxOpenConns(100)              //最多可容纳
	sqlDB.SetConnMaxLifetime(time.Hour * 4) //连接最大复用时间， 不能超过mysql的wait_timeout
	return db
}

func MysqlConnect() {

}
