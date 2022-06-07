package common

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/app/controller/base"
	"mqenergy-go/config"
	"mqenergy-go/pkg/auth"
	"mqenergy-go/pkg/response"
	"net/http"
	"strings"
)

type TokenController struct {
	base.Controller
}

var Token = TokenController{}

// Create 生成token
func (c *TokenController) Create(ctx *gin.Context) {
	token, err := auth.GenerateJwtToken(config.Conf.Jwt.Secret, config.Conf.Jwt.TokenExpire, map[string]interface{}{"id": 1}, config.Conf.Jwt.TokenIssuer)
	if err != nil {
		response.UnauthorizedException(ctx, err.Error())
		return
	}
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", token)
}

// View token解析
func (c *TokenController) View(ctx *gin.Context) {
	token := ctx.GetHeader(config.Conf.Jwt.TokenKey)
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
	jwtTokenArr, err := auth.ParseJwtToken(token, config.Conf.Jwt.Secret)
	if err != nil {
		response.UnauthorizedException(ctx, "")
		return
	}
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", jwtTokenArr)
}
