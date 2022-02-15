package controller

import (
	"github.com/gin-gonic/gin"
	"lyky-go/app/service/common"
	"lyky-go/entities"
	"lyky-go/pkg/response"
)

type BaseController struct{}

var Base = BaseController{}

// Ping
func (c BaseController) Ping(ctx *gin.Context) {
	response.OkWithData(ctx, "", "Pong!")
}

// Login 用户登录
func (c BaseController) Login(ctx *gin.Context) {
	var params entities.UserLoginRequest
	err := ctx.ShouldBindJSON(&params)
	if err != nil {
		response.Response(ctx, response.RequestParamErr, response.CodeMap[response.RequestParamErr], interface{}(nil))
		return
	}
	token, err := common.User.Login(params)
	if err != nil {
		response.FailWithData(ctx, err.Error(), "")
		return
	}
	response.OkWithData(ctx, response.CodeMap[response.Success], token)
}
