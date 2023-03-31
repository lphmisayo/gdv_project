package models

import "gvd_project/models/ctype"

type ArticleModel struct {
	MODEL
	Title         string         `json:"title"  gorm:"size:32"`
	Abstract      string         `json:"abstract"`
	Content       string         `json:"content"`
	ViewCount     int            `json:"viewCount"`
	CommentCount  int            `json:"comment_count"`
	DiggCount     int            `json:"digg_count"`
	CollectCount  int            `json:"collect_count"`
	TagModels     []TagModel     `json:"tag_models" gorm:"many2many:article_tag"`
	CommentModels []CommentModel `json:"-" gorm:"foreignKey:ArticleID"`
	UserModel     UserModel      `json:"-" gorm:"foreignKey:UserID"`
	UserID        uint           `json:"user_id"`
	Category      string         `json:"category" gorm:"size:20"`
	Source        string         `json:"source"`
	Link          string         `json:"link"`
	Banner        BannerModel    `json:"-" gorm:"foreignKey:BannerID"`
	BannerID      uint           `json:"banner_id"`
	NickName      string         `json:"nick_name" gorm:"size:42"`
	BannerPath    string         `json:"banner_path"`
	Tags          ctype.Array    `json:"tags" gorm:"type:string;size:64"`
}
