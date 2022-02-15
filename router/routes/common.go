package routes

import (
	"github.com/gin-gonic/gin"
	"lyky-go/app/controller"
	"lyky-go/app/controller/common"
)

func InitPublicCommonRouter(r *gin.RouterGroup) (router gin.IRoutes) {
	commonGroup := r.Group("")
	{
		// ping
		commonGroup.GET("/ping", controller.Base.Ping)
		// 生成token
		commonGroup.GET("/token-create", common.Token.TokenCreate)
		// 解析token
		commonGroup.POST("/token-parse", common.Token.TokenParse)
		// 登录
		commonGroup.POST("/user/login", controller.Base.Login)
	}
	return commonGroup
}
