package middleware

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mqenergy-go/global"
	"time"
)

type CustomResponseWriter struct {
	gin.ResponseWriter
	body *bytes.Buffer
}

func (w CustomResponseWriter) Write(b []byte) (int, error) {
	w.body.Write(b)
	return w.ResponseWriter.Write(b)
}

func (w CustomResponseWriter) WriteString(s string) (int, error) {
	w.body.WriteString(s)
	return w.ResponseWriter.WriteString(s)
}

// RequestIdAuth requestId中间件
func RequestIdAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		writer := CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = writer
		// 开始时间
		reqStartTime := time.Now().UnixMilli()
		// 处理请求
		ctx.Next()
		if ctx.Request.Method != "GET" {
			ctx.Request.ParseForm()
		}
		var responseData string
		// 结束时间
		reqEndTime := time.Now().UnixMilli()
		fields := logrus.Fields{
			"req_body":     ctx.Request.PostForm.Encode(),
			"req_host":     ctx.Request.Host,
			"req_method":   ctx.Request.Method,
			"req_clientIp": ctx.ClientIP(),
			"req_uri":      ctx.Request.RequestURI,
			"res_time":     fmt.Sprintf("%vms", reqEndTime-reqStartTime), // 响应时间
		}
		if ctx.Writer.Status() != 200 {
			responseData = writer.body.String()
			fields["req_header"] = ctx.Request.Header
			// 记录request日志
			global.Logger.WithFields(fields).Warn(responseData)
		} else {
			global.Logger.WithFields(fields).Info(responseData)
		}
	}
}
