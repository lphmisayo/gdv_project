package common

import (
	"gorm.io/gorm"
	"gvd_server/global"
	"gvd_server/models"
)

type Option struct {
	models.PageInfo
	Debug bool
}

func CommonList[T any](model T, option Option) (list []T, count int64, err error) {

	//设置mysql日志
	DB := global.DB
	if option.Debug {
		DB = global.DB.Session(&gorm.Session{Logger: global.MysqlLogConfig.MysqlLog})
	}

	//定制排序默认值
	if option.Sort == "" {
		option.Sort = "create_at desc" //默认时间倒序
	}

	count = DB.Select("id").Find(&list).RowsAffected
	//确认页数
	offest := (option.Page - 1) * option.Limit
	if offest < 0 {
		offest = 0
	}
	err = DB.Limit(option.Limit).Offset(offest).Order(option.Sort).Find(&list).Error
	return list, count, err
}
