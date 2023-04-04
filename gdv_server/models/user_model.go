package models

import "gvd_server/models/ctype"

type UserModel struct {
	MODEL
	UserName       string           `json:"user_name"  gorm:"size:42"`
	NickName       string           `json:"nick_name" gorm:"size:42"`
	PassWord       string           `json:"pass_word" gorm:"size:128"`
	Avatar         string           `json:"avatar" gorm:"size:500"`
	Email          string           `json:"email" gorm:"128"`
	Tel            string           `json:"tel" gorm:"size:18"`
	Addr           string           `json:"addr" gorm:"size:64"`
	Token          string           `json:"token" gorm:"size:64"`
	IP             string           `json:"IP" gorm:"size:20"`
	Role           ctype.Role       `json:"role" gorm:"size:4;default:1"`
	SignStatus     ctype.SignStatus `json:"sign_status" gorm:"type=smallint(6)"`
	ArticleModels  []ArticleModel   `json:"-" gorm:"foreignKey:UserID"`
	CollectsModels []ArticleModel   `json:"-" gorm:"many2many:user_collect_model;joinForeignKey:UserID;JoinReferences:ArticleID"`
}
