package models

import (
	"gorm.io/gorm/logger"
	"time"
)

//存储公共类

type MODEL struct {
	ID        uint      ` gorm:"primaryKey" json:"id"` //主键id
	CreateAt  time.Time ` json:"create_time"`          //创建时间
	UpdatedAt time.Time ` json:"-"`                    //更新时间
}

type PageInfo struct {
	Page  int    `form:"page"`
	Key   string `form:"key"`
	Limit int    `form:"limit"`
	Sort  string `form:"sort"`
}

type MysqlLogSt struct {
	MysqlLog logger.Interface
	Debug    bool
}
