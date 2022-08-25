package common

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/pkg/auth"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/gin-gonic/gin"
	"strings"
)

type TokenController struct {
	base.Controller
}

var Token = TokenController{}

// Create 生成token
func (c *TokenController) Create(ctx *gin.Context) {
	token, err := auth.GenerateJwtToken(global.Cfg.Jwt.Secret, global.Cfg.Jwt.TokenExpire, map[string]interface{}{"id": 1}, global.Cfg.Jwt.TokenIssuer)
	if err != nil {
		response.UnauthorizedException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", token)
}

// View token解析
func (c *TokenController) View(ctx *gin.Context) {
	token := ctx.GetHeader(global.Cfg.Jwt.TokenKey)
	if token == "" {
		response.UnauthorizedException(ctx, "")
		return
	}
	flag := strings.Contains(token, "Bearer")
	if !flag {
		response.UnauthorizedException(ctx, "")
		return
	}
	token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
	jwtTokenArr, err := auth.ParseJwtToken(token, global.Cfg.Jwt.Secret)
	if err != nil {
		response.UnauthorizedException(ctx, "")
		return
	}
	response.SuccessJson(ctx, "", jwtTokenArr)
}
