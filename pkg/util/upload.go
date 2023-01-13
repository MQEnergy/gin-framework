package util

import (
	"errors"
	"fmt"
	"github.com/MQEnergy/gin-framework/global"
	"mime/multipart"
	"strconv"
	"strings"
)

// MaxUploadSize 默认最大上传资源大小是10M
const MaxUploadSize = 10 * 1024 * 1024

// AllowTypes 默认允许上传的文件类型
var AllowTypes = []string{"jpg", "jpeg", "png", "svg", "gif", "bmp", "mp3", "mp4", "avi", "pdf", "xls", "xlsx", "ppt", "doc", "docx"}

type Upload struct {
	MaxUploadSize int
	AllowTypes    []string
}

// FileHeader 文件参数
type FileHeader struct {
	Filename   string `json:"file_name"`   // 图片新名称
	FileSize   int64  `json:"file_size"`   // 图片大小
	FilePath   string `json:"file_path"`   // 相对路径地址
	OriginName string `json:"origin_name"` // 图片原名称
	MimeType   string `json:"mime_type"`   // 附件mime类型
	Extension  string `json:"extension"`   // 附件后缀名
}

// NewUpload
// @Description: 携带上传参数 并且实例化
// @param max
// @param allowTypes
// @return *Upload
func NewUpload(maxSize int, allowTypes []string) *Upload {
	if maxSize == 0 {
		maxSize = MaxUploadSize
	}
	if len(allowTypes) == 0 {
		allowTypes = AllowTypes
	}
	return &Upload{
		maxSize,
		allowTypes,
	}
}

// UploadFile
// @Description: 上传图片
// @receiver u
// @param file 请求的文件
// @param path 子目录名称
// @return *FileHeader
// @return error
func (u *Upload) UploadFile(file *multipart.FileHeader, path string) (*FileHeader, error) {
	fileType := strings.Split(file.Header.Get("Content-Type"), "/")[1]
	fileName := fmt.Sprintf("file-%s.%s", GenerateUuid(32), fileType)

	if int64(u.MaxUploadSize) < file.Size {
		return nil, errors.New("超过最大上传大小 不能超过" + strconv.Itoa(u.MaxUploadSize/(1000*1000)) + "M")
	}
	if !InAnySlice[string](u.AllowTypes, fileType) {
		return nil, errors.New("上传文件格式错误 支持格式" + strings.Join(AllowTypes, ","))
	}
	// 创建时间目录
	filePath, err := MakeTimeFormatDir(global.Cfg.Server.FileUploadPath, path, "2006-01-02")
	if err != nil {
		return nil, err
	}
	// 写入文件
	if err := WriteContentToFile(file, global.Cfg.Server.FileUploadPath+filePath+fileName); err != nil {
		return nil, err
	}
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
