package backend

import (
	"github.com/MQEnergy/go-framework/app/controller/base"
	"github.com/MQEnergy/go-framework/pkg/response"
	"github.com/MQEnergy/go-framework/pkg/util"
	"github.com/MQEnergy/go-framework/types/attachment"
	"github.com/gin-gonic/gin"
	"net/http"
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
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", file)
}
