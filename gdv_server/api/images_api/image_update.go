package images_api

import (
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/models/res"
	"gvd_server/utils"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" bingding:"required" msg:"请输入文件ID"`
	Name string `json:"name" bingding:"required" msg:"请输入文件名"`
}

func (ImageApi) ImageUpdateNameView(c *gin.Context) {
	var updateModel ImageUpdateRequest
	err := c.ShouldBindJSON(&updateModel)
	if err != nil {
		res.FailWithError(err, &updateModel, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, updateModel.ID).Error
	if err != nil {
		res.FailWithMsg("文件不存在", c)
		return
	}
	imageModel.Name = updateModel.Name
	imageModel.UpdatedAt = utils.GetNowTimeHasFormat()
	err = global.DB.Updates(&imageModel).Error
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	res.OKWithMsg("图片名称修改成功！", c)
}
