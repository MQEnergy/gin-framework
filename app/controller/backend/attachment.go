package backend

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/app/controller/base"
	"mqenergy-go/entities/attachment"
	"mqenergy-go/pkg/response"
	"mqenergy-go/pkg/util"
	"net/http"
)

type AttachmentController struct {
	base.Controller
}

var Attachment = AttachmentController{}

// Upload 上传图片
func (c AttachmentController) Upload(ctx *gin.Context) {
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
