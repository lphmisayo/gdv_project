package image_ser

import (
	"fmt"
	"gvd_server/global"
	"gvd_server/models"
	"gvd_server/models/ctype"
	"gvd_server/utils"
	"io"
	"mime/multipart"
	"strings"
)

var (
	//WhiteImageList 图片上传白名单
	WhiteImageList []string = []string{
		"jpg", "png", "jpeg", "ico", "tiff", "gif", "svg", "webp",
	}
)

// ImageUploadService 处理图片文件上传的方法
func (ImageService) ImageUploadService(file *multipart.FileHeader, filePath string) (res FileUploadResponse) {
	//初始化响应体
	fileName := file.Filename
	res.FileName = fileName
	res.Code = -1
	res.Msg = "上传失败，非法文件"

	//文件白名单相应
	nameList := strings.Split(fileName, ".")
	suffix := strings.ToLower(nameList[len(nameList)-1])
	//检查上传文件类型是否匹配
	if !utils.CheckInStrList(suffix, WhiteImageList) {
		res = initFileUploadResModel(-1, file, "上传失败，文件类型不匹配")
		return res
	}
	//判断大小
	fileSize := float64(file.Size) / float64(1024*1024)
	fmt.Println(filePath, fileSize)
	fmt.Println(global.Config.Local.Size)
	if fileSize > float64(global.Config.Local.Size) {
		res = initFileUploadResModel(-1, file, "上传失败：超出上传大小")
		return res
	}

	//生成文件md5
	fileObject, err := file.Open()
	if err != nil {
		global.Log.Error(err.Error())
	}
	byteData, err := io.ReadAll(fileObject)
	if err != nil {
		global.Log.Error(err.Error())
	}
	md5Str := utils.MD5(string(byteData))
	fmt.Println(md5Str)
	var bannerModel models.BannerModel
	errDB := global.DB.Take(&bannerModel, "hash = ?", md5Str).Error
	if errDB == nil {
		// 找到了
		res = initFileUploadResModel(-1, file, "上传失败：该文件已存在！")
		return res
	}

	//图片入库
	bannerModel = models.BannerModel{
		Hash:      md5Str,
		Path:      filePath,
		Name:      fileName,
		ImageType: ctype.Local,
	}
	bannerModel.MODEL = models.MODEL{CreateAt: utils.GetNowTimeHasFormat()}
	global.DB.Create(&bannerModel)

	return initFileUploadResModel(0, file, "上传成功！")
}

func initFileUploadResModel(code int, file *multipart.FileHeader, msg string) (f FileUploadResponse) {
	f = FileUploadResponse{
		FileName: file.Filename,
		Code:     code,
		Msg:      msg,
	}
	return f
}

type FileUploadResponse struct {
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}
