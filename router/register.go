package router

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/config"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/middleware"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/MQEnergy/gin-framework/router/routes"
	"github.com/gin-contrib/cors"
	"github.com/gin-contrib/requestid"
	"github.com/gin-gonic/gin"
)

func Register() *gin.Engine {
	gin.SetMode(global.Cfg.Server.Mode)
	r := gin.New()
	r.Use(gin.Recovery())
	// [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	r.SetTrustedProxies(config.AllowIpList)
	// header add X-Request-Id
	r.Use(requestid.New())
	r.Use(middleware.RequestIdAuth())
	// 404 not found
	r.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		response.NotFoundException(ctx, fmt.Sprintf("%s %s not found", method, path))
	})

	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		commonGroup = r.Group("/", publicMiddleware...)
		authGroup   = r.Group("/", append(publicMiddleware, middleware.LoginAuth(), middleware.CasbinAuth())...)
	)
	// 公用组
	routes.InitCommonGroup(commonGroup)
	// 后台组
	routes.InitBackendGroup(authGroup)
	// 前台组
	routes.InitFrontendGroup(authGroup)
	// 赋给全局
	global.Router = r
	return r
}
