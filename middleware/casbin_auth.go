package middleware

import (
	"github.com/casbin/casbin/v2"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"lyky-go/global"
	"lyky-go/global/app"
	"net/http"
	"strconv"
)

// CasbinAuth 用户权限验证
func CasbinAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adapter, _ := gormadapter.NewAdapterByDB(global.DB)
		e, _ := casbin.NewEnforcer("rbac_model.conf", adapter)
		err := e.LoadPolicy()
		//	获取当前请求的url
		obj := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		user, err := app.GetLoginUser(ctx)
		if err != nil {
			ctx.JSON(http.StatusUnauthorized, "权限异常")
			ctx.Abort()
		}
		var flag = false
		for _, sub := range user.RoleIds {
			//	判断策略中是否存在
			subStr := strconv.FormatUint(sub, 10)
			if ok, _ := e.Enforce(subStr, obj, act); ok {
				flag = true
			}
		}
		if flag {
			ctx.Next()
		} else {
			ctx.JSON(http.StatusOK, "该用户无此权限")
			ctx.Abort()
		}
	}
}
