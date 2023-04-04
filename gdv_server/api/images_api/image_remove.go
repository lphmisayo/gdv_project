package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/models/res"
)

func (i ImageApi) ImageRemoveView(c *gin.Context) {
	var remove models.RemoveRequest
	err := c.ShouldBindJSON(&remove)
	if err != nil {
		res.FailWithCode(res.ParamError, c)
		return
	}

	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, remove.IDList).RowsAffected
	if count <= 0 {
		res.FailWithMsg("文件不存在", c)
		return
	}
	err = global.DB.Delete(&imageList).Error
	fmt.Println(err)
	res.OKWithMsg(fmt.Sprintf("共删除%d张图片", count), c)
}
