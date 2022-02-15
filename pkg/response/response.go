package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type JsonResponse struct {
	Code    Code        `json:"code"`    // 错误码
	Message string      `json:"message"` // 提示信息
	Data    interface{} `json:"data"`    // 返回数据（业务接口定义具体数据结构）
}

// Response 基础返回
func Response(ctx *gin.Context, code Code, message string, data interface{}) {
	ctx.JSON(http.StatusOK, JsonResponse{
		Code:    code,
		Message: message,
		Data:    data,
	})
}

// Ok 正确返回
func Ok(ctx *gin.Context) {
	Response(ctx, Success, CodeMap[Success], interface{}(nil))
}

// OkWithData 失败带数据返回
func OkWithData(ctx *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = CodeMap[Success]
	}
	Response(ctx, Success, msg, data)
}

// Fail 失败返回
func Fail(ctx *gin.Context) {
	Response(ctx, Failed, CodeMap[Failed], interface{}(nil))
}

// FailWithData 失败带数据返回
func FailWithData(ctx *gin.Context, msg string, data interface{}) {
	if msg == "" {
		msg = CodeMap[Failed]
	}
	Response(ctx, Failed, msg, data)
}
