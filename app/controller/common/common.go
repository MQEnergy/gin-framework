package common

import (
	"github.com/MQEnergy/gin-framework/app/controller/base"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/models"
	"github.com/MQEnergy/gin-framework/pkg/response"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"net/http"
	"strings"
)

type CommonController struct {
	*base.Controller
}

var Common = &CommonController{}

type routeInfo struct {
	Method     string   `json:"method"`
	Path       string   `json:"path"`
	NewPath    string   `json:"new_path"`
	MethodList []string `json:"method_list"`
}

// Ping 心跳
func (c *CommonController) Ping(ctx *gin.Context) {
	response.ResponseJson(ctx, http.StatusOK, response.Success, "Pong!", "")
}

// Router 获取所有路由
func (c *CommonController) Routes(ctx *gin.Context) {

	var (
		casbinRule     = make([]models.CasbinRule, 0)
		routeList      = make([]routeInfo, 0)
		newPathList    = make([]string, 0)
		casbinRuleList = make([]models.CasbinRule, 0)
	)
	for _, route := range global.Router.Routes() {
		var _routeInfo = routeInfo{}
		if err := copier.Copy(&_routeInfo, route); err != nil {
			continue
		}
		routerArr := strings.Split(route.Path, "/")[1:]
		if len(routerArr) < 2 {
			_routeInfo.NewPath = "/" + strings.Join(routerArr, "/")
		} else {
			routerArr = strings.Split(route.Path, "/")[1:len(routerArr)]
			_routeInfo.NewPath = "/" + strings.Join(routerArr, "/") + "/*"
		}
		if strings.Contains(_routeInfo.NewPath, "backend") {
			routeList = append(routeList, _routeInfo)
			newPathList = append(newPathList, _routeInfo.NewPath)
		}
	}
	newPathList = lo.Uniq[string](newPathList)
	_routeList := make([]routeInfo, len(newPathList))

	for i := 0; i < len(newPathList); i++ {
		_routeList[i].NewPath = newPathList[i]
		for j := 0; j < len(routeList); j++ {
			if routeList[j].NewPath == newPathList[i] {
				_routeList[i].MethodList = append(_routeList[i].MethodList, routeList[j].Method)
			}
		}
	}
	for i := 0; i < len(_routeList); i++ {
		_routeList[i].MethodList = lo.Uniq[string](_routeList[i].MethodList)
		casbinRule = append(casbinRule, models.CasbinRule{
			Ptype: "p",
			V0:    "1",
			V1:    _routeList[i].NewPath,
			V2:    strings.Join(_routeList[i].MethodList, ","),
		})
	}
	// 存入到 casbin_rule中
	for _, rule := range casbinRule {
		var casbinRuleInfo models.CasbinRule
		if err := global.DB.Where("v0 = 1 and v1 = ?", rule.V1).First(&casbinRuleInfo).Error; err != nil {
			casbinRuleList = append(casbinRuleList, rule)
		}
	}
	if len(casbinRuleList) > 0 {
		if err := global.DB.Model(models.CasbinRule{}).Create(casbinRuleList).Error; err != nil {
			response.BadRequestException(ctx, err.Error())
			return
		}
	}
	response.ResponseJson(ctx, http.StatusOK, response.Success, "", casbinRuleList)
}
