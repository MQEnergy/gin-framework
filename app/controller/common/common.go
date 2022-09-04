package common

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
	"github.com/MQEnergy/gin-framework/app/service/common"
	"github.com/MQEnergy/gin-framework/pkg/response"
	common2 "github.com/MQEnergy/gin-framework/types/common"
	"github.com/gin-gonic/gin"
)

type CommonController struct {
	*base.Controller
}

var Common = &CommonController{}

// Ping 心跳
func (c *CommonController) Ping(ctx *gin.Context) {
	response.SuccessJson(ctx, "Pong!", "")
}

// Routes 获取所有路由
func (c *CommonController) Routes(ctx *gin.Context) {
	var requestParams common2.RouteRequest
	if err := c.ValidateReqParams(ctx, &requestParams); err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	casbinRuleList, err := common.Common.AddRoutes(requestParams)
	if err != nil {
		response.BadRequestException(ctx, err.Error())
		return
	}
	response.SuccessJson(ctx, "", casbinRuleList)
}
