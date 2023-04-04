package models

type AdvertModel struct {
	MODEL
	Title  string `json:"title" gorm:"size:32"`
	Href   string `json:"href"`
	Image  string `json:"image"`
	isShow bool   `json:"isShow"`
}
