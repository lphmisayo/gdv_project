package models

import "time"

type User2Collect struct {
	UserID       uint         `gorm:"primaryKey"`
	UserModel    UserModel    `gorm:"ForeignKey:UserID"`
	ArticleID    uint         `gorm:"primaryKey"`
	ArticleModel ArticleModel `gorm:"ForeignKey:ArticleID"`
	CreateAt     time.Time
}
