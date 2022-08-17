package middleware

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"strings"
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
		// 开始时间
		reqStartTime := time.Now()
		writer := CustomResponseWriter{
			body:           bytes.NewBufferString(""),
			ResponseWriter: ctx.Writer,
		}
		ctx.Writer = writer
		// 获取请求参数
		reqBody := getRequestParams(ctx)
		// 处理请求
		ctx.Next()
		fields := logrus.Fields{
			"req_body":     reqBody,
			"req_host":     ctx.Request.Host,
			"req_method":   ctx.Request.Method,
			"req_clientIp": ctx.ClientIP(),
			"req_id":       requestid.Get(ctx),
			"req_uri":      ctx.Request.RequestURI,
			"res_time":     fmt.Sprintf("%s", time.Now().Sub(reqStartTime)), // 响应时间
		}

		if ctx.Writer.Status() != 200 {
			responseData := writer.body.String()
			fields["req_header"] = ctx.Request.Header
			// 记录request日志
			global.Logger.WithFields(fields).Warn(responseData)
		} else {
			global.Logger.WithFields(fields).Info("")
		}
	}
}

// getRequestParams 获取请求参数（GET,POST,PUT,DELETE）等
func getRequestParams(ctx *gin.Context) string {
	if ctx.Request.Method == "GET" {
		var params []string
		values := ctx.Request.URL.Query()
		for key, value := range values {
			params = append(params, key+"="+value[0])
		}
		return strings.Join(params, "&")
	}
	if ctx.ContentType() != "application/json" {
		ctx.Request.ParseForm()
		return ctx.Request.PostForm.Encode()
	}
	rawData, _ := ctx.GetRawData()
	//读取后，重新赋值 c.Request.Body ，以供后续的其他操作
	ctx.Request.Body = ioutil.NopCloser(bytes.NewBuffer(rawData))
	var m map[string]string
	var params []string
	// 反序列化
	json.Unmarshal(rawData, &m)
	for key, value := range m {
		params = append(params, key+"="+value)
	}
	return strings.Join(params, "&")
}
