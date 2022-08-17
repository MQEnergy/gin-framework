package middleware

import (
	"github.com/MQEnergy/go-framework/global"
	"github.com/MQEnergy/go-framework/pkg/auth"
	"github.com/MQEnergy/go-framework/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

// LoginAuth 登录中间件
func LoginAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.GetHeader(global.Cfg.Jwt.TokenKey)
		if token == "" {
			response.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		b := strings.Contains(token, "Bearer")
		if !b {
			response.UnauthorizedException(ctx, "")
			ctx.Abort()
			return
		}
		token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
		if _, err := auth.ParseJwtToken(token, global.Cfg.Jwt.Secret); err != nil {
			response.UnauthorizedException(ctx, err.Error())
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
