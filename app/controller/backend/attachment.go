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
	var requestParams attachment.UploadRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	file, err := util.UploadFile(requestParams.FilePath, ctx)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", file)
}
