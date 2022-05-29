package common

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/app/controller/base"
	"mqenergy-go/config"
	"mqenergy-go/global"
	"mqenergy-go/models"
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
	var adminInfo models.GinAdmin
	global.DB.Where("id = 1").First(&adminInfo)
	token, err := auth.GenerateJwtToken(config.Conf.Server.JwtSecret,
		config.Conf.Server.TokenExpire, adminInfo,
		config.Conf.Server.TokenIssuer)
	if err != nil {
		response.UnauthorizedException(ctx, err.Error())
		return
	}
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", token)
}

// View token解析
func (c *TokenController) View(ctx *gin.Context) {
	token := ctx.GetHeader(config.Conf.Server.TokenKey)
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
	jwtTokenArr, err := auth.ParseJwtToken(token, config.Conf.Server.JwtSecret)
	if err != nil {
		response.UnauthorizedException(ctx, "")
		return
	}
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", jwtTokenArr)
}
