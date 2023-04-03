package router

import (
	"github.com/gin-gonic/gin"
	"gvd_project/api"
)

func ImagesRouter(router *gin.RouterGroup) {
	imageApi := api.ApiGroupApp.ImageApi
	{
		router.POST("image", imageApi.ImageUploadView)
		router.POST("images", imageApi.ImagesUploadView)
	}
}
