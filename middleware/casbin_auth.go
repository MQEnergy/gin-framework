package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"mqenergy-go/global"
	"mqenergy-go/global/app"
	"mqenergy-go/pkg/response"
	"os"
	"strconv"
	"strings"
)

// CasbinAuth 用户权限验证
func CasbinAuth() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		adapter, _ := gormadapter.NewAdapterByDB(global.DB)
		dir, err := os.Getwd()
		if err != nil {
			ctx.Abort()
			return
		}
		e, _ := casbin.NewEnforcer(dir+"/config/rbac_model.conf", adapter)
		e.AddFunction("ParamsMatch", ParamsMatchFunc)
		_ = e.LoadPolicy()

		//	获取当前请求的url
		obj := ctx.Request.URL.RequestURI()
		act := ctx.Request.Method
		user, err := app.GetAdminInfo(ctx)
		if err != nil {
			response.UnauthorizedException(ctx, "权限异常")
			ctx.Abort()
			return
		}
		var flag = false
		for _, sub := range user.RoleIds {
			//	判断策略中是否存在
			subStr := strconv.FormatUint(sub, 10)
			if ok, _ := e.Enforce(subStr, obj, act); ok {
				flag = true
				break
			}
		}
		if !flag {
			response.UnauthorizedException(ctx, "该用户无此权限")
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	return ParamsMatch(name1, name2), nil
}

// ParamsMatch 自定义规则函数
func ParamsMatch(fullNameKey1 string, key2 string) bool {
	key1 := strings.Split(fullNameKey1, "?")[0]
	return util.KeyMatch2(key1, key2)
}
