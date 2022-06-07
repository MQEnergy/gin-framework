package util

import (
	"errors"
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"mqenergy-go/config"
	"os"
	"strconv"
	"strings"
	"time"
)

// 最大上传资源大小是10M
var (
	MaxUploadSize = 10 * 1024 * 1024
	AllowTypes    = []string{"jpg", "jpeg", "png", "svg", "gif", "bmp", "mp3", "mp4", "avi", "pdf", "xls", "xlsx", "ppt", "doc", "docx"}
)

type FileHeader struct {
	Filename   string `json:"file_name"`   // 图片新名称
	FileSize   int64  `json:"file_size"`   // 图片大小
	FilePath   string `json:"file_path"`   // 相对路径地址
	OriginName string `json:"origin_name"` // 图片原名称
	MimeType   string `json:"mime_type"`   // 附件mime类型
	Extension  string `json:"extension"`   // 附件后缀名
}

// UploadFile 上传图片
func UploadFile(path string, r *gin.Context) (*FileHeader, error) {
	file, err := r.FormFile("file")
	if err != nil {
		return nil, err
	}
	if int64(MaxUploadSize) < file.Size {
		return nil, errors.New("超过最大上传大小 不能超过" + strconv.Itoa(MaxUploadSize/(1000*1000)) + "M")
	}
	fileType := strings.Split(file.Header.Get("Content-Type"), "/")[1]
	if !InAnySlice[string](AllowTypes, fileType) {
		return nil, errors.New("上传文件格式错误 支持格式" + strings.Join(AllowTypes, ","))
	}
	fileName := fmt.Sprintf("file-%s.%s", GenerateUuid(32), fileType)
	filePath := "upload/"
	if path != "" {
		filePath += path + "/"
	}
	filePath += time.Now().Format("2006-01-02") + "/"
	b := MakeMultiDir(config.Conf.Server.FileUploadPath + filePath)
	if b != nil {
		return nil, err
	}
	create, err := os.Create(config.Conf.Server.FileUploadPath + filePath + fileName)
	if err != nil {
		return nil, err
	}
	defer create.Close()
	open, err := file.Open()
	if err != nil {
		return nil, err
	}
	fileBytes, err := ioutil.ReadAll(open)
	if err != nil {
		return nil, err
	}
	create.Write(fileBytes)
	fileHeader := FileHeader{
		Filename:   fileName,
		FileSize:   file.Size,
		FilePath:   filePath + fileName,
		OriginName: file.Filename,
		MimeType:   file.Header.Get("Content-Type"),
		Extension:  fileType,
	}
	return &fileHeader, nil
}
