package middleware

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/global"
	"mqenergy-go/pkg/auth"
	"mqenergy-go/pkg/response"
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
