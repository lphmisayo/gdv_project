package models

import "gvd_project/models/ctype"

type UserModel struct {
	MODEL
	UserName       string           `json:"user_name"  grom:"size:42"`
	NickName       string           `json:"nick_name" grom:"size:42"`
	PassWord       string           `json:"pass_word" grom:"size:128"`
	Avatar         string           `json:"avatar" grom:"size:500"`
	Email          string           `json:"email" grom:"128"`
	Tel            string           `json:"tel" grom:"size:18"`
	Addr           string           `json:"addr" grom:"size:64"`
	Token          string           `json:"token" grom:"size:64"`
	IP             string           `json:"IP" grom:"size:20"`
	Role           ctype.Role       `json:"role" grom:"size:4;default:1"`
	SignStatus     ctype.SignStatus `json:"sign_status" grom:"type=smallint(6)"`
	ArticleModels  []ArticleModel   `json:"-" grom:"foreignKey:user_id"`
	CollectsModels []ArticleModel   `json:"-" grom:"many2many:auth2_collect;joinForeignKey:user_id;joinReference:article_id"`
}
