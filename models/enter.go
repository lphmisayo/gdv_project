package models

import "time"

type MODEL struct {
	ID        uint      ` gorm:"primaryKey" json:"id"` //主键id
	CreateAt  time.Time ` json:"create_time"`          //创建时间
	UpdatedAt time.Time ` json:"-"`                    //更新时间
}
