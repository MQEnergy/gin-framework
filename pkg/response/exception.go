package response

import (
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	Status    int         `json:"status"`
	ErrCode   Code        `json:"errcode"`
	RequestId string      `json:"requestid"`
	Message   string      `json:"message"`
	Data      interface{} `json:"data"`
}

// ResponseJson 基础返回
func ResponseJson(ctx *gin.Context, status int, errcode Code, message string, data interface{}) {
	if message == "" {
		message = CodeMap[errcode]
	}

	ctx.JSON(status, JsonResponse{
		Status:    status,
		ErrCode:   errcode,
		Message:   message,
		RequestId: requestid.Get(ctx),
		Data:      data,
	})
}

// BadRequestException 400错误
func BadRequestException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[RequestParamErr]
	}
	ResponseJson(ctx, http.StatusBadRequest, RequestParamErr, message, nil)
}

// UnauthorizedException 401错误
func UnauthorizedException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[UnAuthed]
	}
	ResponseJson(ctx, http.StatusUnauthorized, UnAuthed, message, nil)
}

// ForbiddenException 403错误
func ForbiddenException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[Failed]
	}
	ResponseJson(ctx, http.StatusForbidden, Failed, message, nil)
}

// NotFoundException 404错误
func NotFoundException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[RequestMethodErr]
	}
	ResponseJson(ctx, http.StatusNotFound, RequestMethodErr, message, nil)
}

// InternalServerException 500错误
func InternalServerException(ctx *gin.Context, message string) {
	if message == "" {
		message = CodeMap[InternalErr]
	}
	ResponseJson(ctx, http.StatusInternalServerError, InternalErr, message, nil)
}
