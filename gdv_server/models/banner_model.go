package models

// BannerModel 文章封面类
type BannerModel struct {
	MODEL
	Path string `json:"path"`         //路径
	Hash string `json:"hash"`         //图片的hash值，用于判断图片是否重复
	Name string `json:"name" gorm:""` //图片名
}
