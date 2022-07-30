package command

import (
	"errors"
	"github.com/sirupsen/logrus"
	"io/ioutil"
	"mqenergy-go/pkg/util"
	"os"
	"strings"
)

func CreateControllerContent(controller string, module string) error {
	moduleName := strings.ToLower(module)
	controllerName := strings.ToLower(controller)
	firstUpperCtlName := strings.ToUpper(controllerName[:1]) + controllerName[1:]
	projectModuleName := util.GetProjectModuleName()
	content := `package ` + moduleName + `

import (
	"github.com/gin-gonic/gin"
	"` + projectModuleName + `/app/controller/base"
)

type ` + firstUpperCtlName + `Controller struct {
	*base.Controller
}

var ` + firstUpperCtlName + ` = &` + firstUpperCtlName + `Controller{}

// Index 获取列表
func (c *` + firstUpperCtlName + `Controller) Index(ctx *gin.Context) {
	// Todo list
}
`
	path := "app/controller/" + module + "/" + controllerName + ".go"
	if flag := util.IsPathExist(path); !flag {
		err := ioutil.WriteFile(path, []byte(content), 0644)
		if err != nil {
			return errors.New(firstUpperCtlName + ".go 创建失败")
		}
		return nil
	}
	return errors.New(controllerName + ".go 已经存在")
}

// GenerateController 创建controller类
func GenerateController() {
	args := os.Args
	if len(args) != 4 {
		logrus.Error("参数错误 请输入参数：controller名称（不限大小写） module名称（不限大小写 如：backend）")
		os.Exit(0)
	}
	controller := args[2]
	module := args[3]
	err := CreateControllerContent(controller, module)
	if err != nil {
		logrus.Error(err)
		os.Exit(2)
	}
	logrus.Info(controller + ".go 创建成功")
}
