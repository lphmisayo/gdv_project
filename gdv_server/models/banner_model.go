package models

import (
	"gorm.io/gorm"
	"gvd_server/models/ctype"
	"os"
)

// BannerModel 文章封面类
type BannerModel struct {
	MODEL
	Path      string          `json:"path"`                        //路径
	Hash      string          `json:"hash"`                        //图片的hash值，用于判断图片是否重复
	Name      string          `json:"name" gorm:""`                //图片名
	ImageType ctype.ImageType `gorm:"default:1" json:"image_type"` //图片类型
}

// BeforeDelete 图片删除钩子函数，在删除数据库数据前先删除本地文件
func (b *BannerModel) BeforeDelete(db *gorm.DB) (err error) {
	if b.ImageType == ctype.Local {
		//如果事本地图片，删除，还要删除本地文件
		err := os.Remove(b.Path)
		if err != nil {
			//global.Log.Error(err.Error())
			return err
		}
	}
	return nil
}
