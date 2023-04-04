package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_server/global"
	"gvd_server/models/res"
	"gvd_server/service"
	"gvd_server/service/image_ser"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
)

var pathDir = "uploads"

/*type FileUploadResponse struct {
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}
*/
func initFileUploadRes(f []image_ser.FileUploadResponse, code int, file *multipart.FileHeader, msg string) []image_ser.FileUploadResponse {
	if f == nil {
		f = make([]image_ser.FileUploadResponse, 0)
	}
	f = append(f, image_ser.FileUploadResponse{
		FileName: file.Filename,
		Code:     code,
		Msg:      msg,
	})
	return f
}

// ImagesUploadView 多个图片上传
func (ImageApi) ImagesUploadView(c *gin.Context) {
	images, err := c.MultipartForm()
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fileList, ok := images.File["image"]
	if !ok {
		res.FailWithCode(res.FileExistError, c)
		return
	}

	//判断路径是否存在
	//不存在就创建
	dir, err := os.ReadDir(global.Config.Local.Path)
	if err != nil {
		err = os.MkdirAll(global.Config.Local.Path, fs.ModePerm)
		fmt.Println(err)
		fmt.Println(dir, err)
	}

	//上传失败数组
	UploadFiles := make([]image_ser.FileUploadResponse, 0)

	for _, file := range fileList {
		filePath := path.Join(global.Config.Local.Path, file.Filename)

		fileResponse := service.ServiceApp.ImageService.ImageUploadService(file, filePath)
		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err.Error())
			fileResponse = initFileUploadResSingle(-1, file, "上传失败："+err.Error())
			UploadFiles = append(UploadFiles, fileResponse)
			continue
		}
		UploadFiles = append(UploadFiles, fileResponse)

	}

	res.OKWithData(UploadFiles, c)
}

func initFileUploadResSingle(code int, file *multipart.FileHeader, msg string) (f image_ser.FileUploadResponse) {
	f = image_ser.FileUploadResponse{
		FileName: file.Filename,
		Code:     code,
		Msg:      msg,
	}
	return f
}