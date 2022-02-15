package routes

import (
	"github.com/gin-gonic/gin"
	"lyky-go/app/controller/backend"
)

// InitBackendGroup 初始化后台接口路由
func InitBackendGroup(r *gin.RouterGroup) gin.IRoutes {
	backendGroup := r.Group("api")
	{
		backendGroup.POST("/user/create", backend.User.Create)
		backendGroup.GET("/user/view", backend.User.View)
		backendGroup.POST("/user/update", backend.User.Update)
		backendGroup.POST("/user/delete", backend.User.Delete)
	}
	return backendGroup
}
