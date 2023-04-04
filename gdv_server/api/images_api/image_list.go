package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/models/res"
	"gvd_server/service/common"
)

func (a ImageApi) ImageListView(c *gin.Context) {
	var pageInit models.PageInfo
	err := c.ShouldBindQuery(&pageInit)
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fmt.Println(pageInit)

	imageList, count, err := common.CommonList(models.BannerModel{}, common.Option{
		PageInfo: pageInit,
		Debug:    global.MysqlLogConfig.Debug,
	})

	res.OKWithList(imageList, count, c)
}
