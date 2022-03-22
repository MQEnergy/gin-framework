package middleware

import (
	"bytes"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"mqenergy-go/global"
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
		body, _ := ctx.GetRawData()
		writer := CustomResponseWriter{body: bytes.NewBufferString(""), ResponseWriter: ctx.Writer}
		ctx.Writer = writer
		ctx.Next()
		var responseData string
		fields := logrus.Fields{
			"_body":     string(body),
			"_host":     ctx.Request.Host,
			"_method":   ctx.Request.Method,
			"_clientIp": ctx.ClientIP(),
			"_uri":      ctx.Request.RequestURI,
		}
		if ctx.Writer.Status() != 200 {
			responseData = writer.body.String()
			fields["_header"] = ctx.Request.Header
			// 记录request日志
			global.Logger.WithFields(fields).Warn(responseData)
		} else {
			global.Logger.WithFields(fields).Info(responseData)
		}
	}
}
