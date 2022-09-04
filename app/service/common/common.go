package common

import (
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/models"
	"github.com/MQEnergy/gin-framework/pkg/util"
	"github.com/MQEnergy/gin-framework/types/common"
	"github.com/jinzhu/copier"
	"github.com/samber/lo"
	"strconv"
	"strings"
)

type CommonService struct{}

var Common = &CommonService{}

// AddRoutes 添加路由权限
func (s *CommonService) AddRoutes(requestParams common.RouteRequest) ([]models.CasbinRule, error) {
	var (
		casbinRule     = make([]models.CasbinRule, 0)
		routeList      = make([]common.RouteInfo, 0)
		newPathList    = make([]string, 0)
		casbinRuleList = make([]models.CasbinRule, 0)
	)
	for _, route := range global.Router.Routes() {
		var _routeInfo = common.RouteInfo{}
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
		if strings.Contains(_routeInfo.NewPath, requestParams.RType) {
			routeList = append(routeList, _routeInfo)
			newPathList = append(newPathList, _routeInfo.NewPath)
		}
	}
	newPathList = lo.Uniq[string](newPathList)
	_routeList := make([]common.RouteInfo, len(newPathList))

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
			V0:    strconv.Itoa(requestParams.RoleId),
			V1:    _routeList[i].NewPath,
			V2:    strings.Join(_routeList[i].MethodList, ","),
		})
	}
	// 存入到 casbin_rule中
	for _, rule := range casbinRule {
		var casbinRuleInfo models.CasbinRule
		if err := global.DB.Where("v0 = ? and v1 = ?", requestParams.RoleId, rule.V1).First(&casbinRuleInfo).Error; err != nil {
			casbinRuleList = append(casbinRuleList, rule)
			continue
		}
		// 存在
		flag := true
		ruleV2 := strings.Split(rule.V2, ",")
		casbinRuleInfoV2 := strings.Split(casbinRuleInfo.V2, ",")
		for _, s2 := range ruleV2 {
			if !util.InAnySlice(casbinRuleInfoV2, s2) {
				// 不存在
				flag = false
				break
			}
		}
		if !flag {
			// 删除记录
			if err := global.DB.Where("v0 = ? and v1 = ?", requestParams.RoleId, rule.V1).Delete(&casbinRuleInfo).Error; err != nil {
				continue
			}
			casbinRuleList = append(casbinRuleList, rule)
		}
	}
	if len(casbinRuleList) > 0 {
		if err := global.DB.Model(models.CasbinRule{}).Create(casbinRuleList).Error; err != nil {
			return nil, err
		}
	}
	return casbinRuleList, nil
}
