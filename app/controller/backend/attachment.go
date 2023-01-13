package backend

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/MQEnergy/gin-framework/pkg/util"
	"github.com/MQEnergy/gin-framework/types/attachment"
	"github.com/gin-gonic/gin"
)

type AttachmentController struct {
	base.Controller
}

var Attachment = AttachmentController{}

// Upload 上传图片
func (c *AttachmentController) Upload(ctx *gin.Context) {
	var reqParams attachment.UploadRequest
	if err := c.ValidateReqParams(ctx, &reqParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	fileHeader, err := util.NewUpload(0, nil).UploadFile(reqParams.FileName, reqParams.FilePath)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", fileHeader)
}
