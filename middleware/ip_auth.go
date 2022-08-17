package middleware

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/config"
	"github.com/gin-gonic/gin"
	"net/http"
)

// IpAuth 白名单验证
func IpAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		clientIp := ctx.ClientIP()
		flag := false
		for _, value := range config.AllowIpList {
			if value == "*" || clientIp == value {
				flag = true
				break
			}
		}
		if !flag {
			ctx.JSON(http.StatusUnauthorized, fmt.Sprintf("%s 不在ip白名单中", clientIp))
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
