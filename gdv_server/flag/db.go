package flag

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
)

func MakeMigration() {
	fmt.Println("这是迁移方法")
	global.DB.SetupJoinTable(&models.UserModel{}, "CollectsModels", &models.UserCollectModel{})
	global.DB.SetupJoinTable(&models.MenuModel{}, "Banners", &models.MenuBannerModel{})
	err := global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.UserModel{},
			&models.BannerModel{},
			&models.ArticleModel{},
			&models.AdvertModel{},
			&models.CommentModel{},
			&models.MenuModel{},
			&models.FadeBackModel{},
			&models.TagModel{},
			&models.MessageModel{},
			&models.LoginDataModel{},
			&models.MenuBannerModel{},
		)
	if err != nil {
		global.Log.Error(err.Error())
	}
}
