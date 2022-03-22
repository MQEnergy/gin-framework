package command

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mqenergy-go/pkg/util"
	"os"
	"strings"
)

func CreateServiceContent(service string, module string) error {
	moduleName := strings.ToLower(module)
	serviceName := strings.ToLower(service)
	firstUpperCtlName := strings.ToUpper(serviceName[:1]) + serviceName[1:]
	content := `package ` + moduleName + `

type ` + firstUpperCtlName + `Service struct{}

var ` + firstUpperCtlName + ` = ` + firstUpperCtlName + `Service{}

// GetList 获取列表
func (s ` + firstUpperCtlName + `Service) GetList(requestParams interface{}) (interface{}, error) {
	// Todo list 
	return nil, nil
}
`
	path := "app/service/" + module + "/" + serviceName + ".go"
	if flag := util.IsPathExist(path); !flag {
		err := ioutil.WriteFile(path, []byte(content), 0644)
		if err != nil {
			return errors.New(firstUpperCtlName + ".go 创建失败")
		}
		return nil
	}
	return errors.New(serviceName + ".go 已经存在")
}

// GenerateService 创建service类
func GenerateService() {
	args := os.Args
	if len(args) != 4 {
		logrus.Error("参数错误 请输入 service名称 module名称")
		os.Exit(0)
	}
	service := args[2]
	module := args[3]
	err := CreateServiceContent(service, module)
	if err != nil {
		logrus.Error(err)
		os.Exit(2)
	}
	logrus.Info(service + ".go 创建成功")
}
