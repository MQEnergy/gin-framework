package backend

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
	"github.com/MQEnergy/gin-framework/app/service/backend"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/MQEnergy/gin-framework/types/user"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	base.Controller
}

var User = UserController{}

// Index 获取列表
func (c *UserController) Index(ctx *gin.Context) {
	var requestParams user.IndexRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	list, err := backend.User.GetIndex(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", list)
}

// List 获取列表
func (c *UserController) List(ctx *gin.Context) {
	var requestParams user.IndexRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	list, err := backend.User.GetList(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", list)
}
