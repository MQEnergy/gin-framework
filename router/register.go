package router

import (
	"fmt"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"lyky-go/config"
	"lyky-go/middleware"
	"lyky-go/router/routes"
	"net/http"
)

func Register() *gin.Engine {
	gin.SetMode(config.Conf.Mode)
	router := gin.New()
	// [WARNING] You trusted all proxies, this is NOT safe. We recommend you to set a value.
	router.SetTrustedProxies(config.AllowIpList)

	//	404处理
	router.NoRoute(func(ctx *gin.Context) {
		path := ctx.Request.URL.Path
		method := ctx.Request.Method
		ctx.JSON(http.StatusNotFound, fmt.Sprintf("%s %s not found", method, path))
	})

	// 路由分组
	var (
		publicMiddleware = []gin.HandlerFunc{
			cors.Default(),
			middleware.IpAuth(),
		}
		commonGroup = router.Group("/", publicMiddleware...)
		authGroup   = router.Group("/", append(publicMiddleware, middleware.LoginAuth, middleware.CasbinAuth())...)
	)
	// 公用组
	routes.InitPublicCommonRouter(commonGroup)
	// 后台组
	routes.InitBackendGroup(authGroup)
	// 前台组
	routes.InitFrontendGroup(authGroup)

	return router
}
