package backend

import (
	"github.com/gin-gonic/gin"
	"mqenergy-go/app/controller/base"
	"mqenergy-go/app/service/common"
	"mqenergy-go/pkg/response"
	"mqenergy-go/types/user"
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
