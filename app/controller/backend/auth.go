package backend

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
	"github.com/MQEnergy/gin-framework/app/service/common"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/MQEnergy/gin-framework/types/user"
	"github.com/gin-gonic/gin"
)

type AuthController struct {
	base.Controller
}

var Auth = AuthController{}

// Login 用户登录
func (c *AuthController) Login(ctx *gin.Context) {
	var requestParams user.LoginRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	token, err := common.User.Login(requestParams)
	if err != nil {
		response.BadRequestException(ctx, "")
		return
	}
	response.SuccessJson(ctx, "", token)
}
