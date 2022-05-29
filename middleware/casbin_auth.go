package middleware

import (
	"github.com/casbin/casbin/v2"
	"github.com/casbin/casbin/v2/util"
	gormadapter "github.com/casbin/gorm-adapter/v3"
	"github.com/gin-gonic/gin"
	"github.com/gogf/gf/v2/util/gconv"
	"mqenergy-go/global"
	"mqenergy-go/global/app"
	"mqenergy-go/pkg/response"
	util2 "mqenergy-go/pkg/util"
	"os"
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
		e.AddFunction("ParamsActMatch", ParamsActMatchFunc)
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
			subStr := gconv.String(sub)
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

// ParamsActMatchFunc 自定义规则函数
func ParamsActMatchFunc(args ...interface{}) (interface{}, error) {
	rAct := args[0].(string)
	pAct := args[1].(string)
	pActArr := strings.Split(pAct, ",")
	return util2.InStringSlice[string](pActArr, rAct), nil
}

// ParamsMatchFunc 自定义规则函数
func ParamsMatchFunc(args ...interface{}) (interface{}, error) {
	name1 := args[0].(string)
	name2 := args[1].(string)
	key1 := strings.Split(name1, "?")[0]
	// 剥离路径后再使用casbin的keyMatch2
	return util.KeyMatch2(key1, name2), nil
}
