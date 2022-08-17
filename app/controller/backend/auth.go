package backend

import (
	"github.com/MQEnergy/go-framework/app/controller/base"
	"github.com/MQEnergy/go-framework/app/service/common"
	"github.com/MQEnergy/go-framework/pkg/response"
	"github.com/MQEnergy/go-framework/types/user"
	"github.com/gin-gonic/gin"
	"net/http"
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
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", token)
}
