package images_api

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"gvd_project/global"
	"gvd_project/models"
	"gvd_project/models/res"
	"gvd_project/utils"
	"io"
	"io/fs"
	"mime/multipart"
	"os"
	"path"
	"strings"
)

var (
	//图片上传白名单
	WhiteImageList []string = []string{
		"jpg", "png", "jpeg", "ico", "tiff", "gif", "svg", "webp",
	}
)

var pathDir = "uploads"

type FileUploadResponse struct {
	FileName string `json:"fileName"`
	Code     int    `json:"code"`
	Msg      string `json:"msg"`
}

// ImageUploadView 单个图片上传
func (i ImageApi) ImageUploadView(c *gin.Context) {
	fileHeader, err := c.FormFile("image")
	if err != nil {
		res.FailWithMsg(err.Error(), c)
		return
	}
	fmt.Println(fileHeader.Header)
	fmt.Println(fileHeader.Size)
	fmt.Println(fileHeader.Filename)
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
	UploadFiles := make([]FileUploadResponse, 0)

	for _, file := range fileList {
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1])
		//检查上传文件类型是否匹配
		if !utils.CheckInStrList(suffix, WhiteImageList) {
			UploadFiles = initFileUploadRes(UploadFiles, -1, file, "上传失败，文件类型不匹配")
			continue
		}
		filePath := path.Join(global.Config.Local.Path, file.Filename)
		//判断大小
		fileSize := float64(file.Size) / float64(1024*1024)
		fmt.Println(filePath, fileSize)
		fmt.Println(global.Config.Local.Size)
		if fileSize > float64(global.Config.Local.Size) {
			UploadFiles = initFileUploadRes(UploadFiles, -1, file, "上传失败：超出上传大小")
			continue
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
			UploadFiles = initFileUploadRes(UploadFiles, -1, file, "上传失败：该文件已存在！")
			continue
		} else {
			fmt.Println(errDB.Error())
		}

		err = c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err.Error())
			UploadFiles = initFileUploadRes(UploadFiles, -1, file, "上传失败："+err.Error())
			continue
		}
		UploadFiles = initFileUploadRes(UploadFiles, 0, file, "上传成功！")
		bannerModel = models.BannerModel{Hash: md5Str, Path: filePath, Name: fileName}
		bannerModel.MODEL = models.MODEL{CreateAt: utils.GetNowTimeHasFormat()}
		global.DB.Create(&bannerModel)

	}

	res.OKWithData(UploadFiles, c)
}

func initFileUploadRes(f []FileUploadResponse, code int, file *multipart.FileHeader, msg string) []FileUploadResponse {
	if f == nil {
		f = make([]FileUploadResponse, 0)
	}
	f = append(f, FileUploadResponse{
		FileName: file.Filename,
		Code:     code,
		Msg:      msg,
	})
	return f
}
