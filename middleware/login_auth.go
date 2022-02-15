package middleware

import (
	"github.com/gin-gonic/gin"
	"lyky-go/config"
	"lyky-go/pkg/auth"
	"lyky-go/pkg/response"
	"net/http"
	"strings"
)

// LoginAuth 登录中间件
func LoginAuth(ctx *gin.Context) {
	token := ctx.GetHeader(config.Conf.Server.TokenKey)
	if token == "" {
		ctx.JSON(http.StatusUnauthorized, response.CodeMap[response.UnAuthed])
		ctx.Abort()
		return
	}
	b := strings.Contains(token, "Bearer")
	if !b {
		ctx.JSON(http.StatusUnauthorized, response.CodeMap[response.UnAuthed])
		ctx.Abort()
		return
	}
	token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
	if _, err := auth.ParseJwtToken(token, config.Conf.Server.JwtSecret); err != nil {
		ctx.Abort()
		return
	}
	ctx.Next()
}
