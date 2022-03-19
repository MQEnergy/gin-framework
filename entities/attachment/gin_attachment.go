package attachment

import "mime/multipart"

type UploadRequest struct {
	FileName *multipart.FileHeader `form:"file" binding:"required"`
	FilePath string                `form:"file_path"`
}
