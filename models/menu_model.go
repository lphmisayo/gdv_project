package models

import "gvd_project/models/ctype"

type MenuModel struct {
	MODEL
	MenuTitle    string        `json:"menu_title" gorm:"size:32"`
	MenuTitleEn  string        `json:"menu_title_en" gorm:"size:32"`
	Slogan       string        `json:"slogan" gorm:"size:64"`
	Abstract     ctype.Array   `json:"abstract" gorm:"type:string"`
	AbstractTime int           `json:"abstract_time"`
	Banners      []BannerModel `gorm:"many2many:menu_banner_models;JoinForeignKey:MenuID;JoinReferences:BannerID" json:"banners"`
	BannerTime   int           `json:"banner_time"`
	Sort         int           `json:"sort" gorm:""`
}
