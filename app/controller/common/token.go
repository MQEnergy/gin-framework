package common

import (
	"github.com/gin-gonic/gin"
	"lyky-go/config"
	"lyky-go/entities"
	"lyky-go/pkg/auth"
	"lyky-go/pkg/response"
	"strings"
)

type TokenController struct{}

var Token = TokenController{}

// TokenCreate 生成token
func (c TokenController) TokenCreate(ctx *gin.Context) {
	var user entities.User
	user.Id = uint64(1)
	token, err := auth.GenerateJwtToken(config.Conf.Server.JwtSecret, config.Conf.Server.TokenExpire, user, config.Conf.Server.TokenIssuer)
	if err != nil {
		response.FailWithData(ctx, err.Error(), "")
		return
	}
	response.OkWithData(ctx, token, "")
}

// TokenParse token解析
func (c TokenController) TokenParse(ctx *gin.Context) {
	token := ctx.GetHeader(config.Conf.Server.TokenKey)
	if token == "" {
		response.FailWithData(ctx, response.CodeMap[response.UnAuthed], "")
		return
	}
	flag := strings.Contains(token, "Bearer")
	if !flag {
		response.FailWithData(ctx, response.CodeMap[response.UnAuthed], "")
		return
	}
	token = strings.TrimSpace(strings.TrimLeft(token, "Bearer"))
	jwtTokenArr, err := auth.ParseJwtToken(token, config.Conf.Server.JwtSecret)
	if err != nil {
		response.FailWithData(ctx, err.Error(), "")
		return
	}
	response.OkWithData(ctx, response.CodeMap[response.Success], jwtTokenArr)
}
