package cmd

import (
	"fmt"
	"github.com/MQEnergy/go-framework/bootstrap"
	"github.com/MQEnergy/go-framework/config"
	"github.com/MQEnergy/go-framework/pkg/util"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"strings"
)

var (
	ctlName string
	modName string
)

// ControllerCmd 数据库模型创建命令
func ControllerCmd() *cli.Command {
	return &cli.Command{
		Name:  "controller",
		Usage: "Create a new controller class",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | test | prod]",
				Destination: &config.ConfEnv,
				Required:    false,
			},
			&cli.StringFlag{
				Name:        "controller",
				Aliases:     []string{"c"},
				Value:       "",
				Usage:       "请输入控制器名称 如：admin",
				Destination: &ctlName,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "module",
				Aliases:     []string{"m"},
				Value:       "backend",
				Usage:       "请输入模块名称 如：backend",
				Destination: &modName,
				Required:    false,
			},
		},
		Action: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.LoggerService)
			return generateController()
		},
	}
}

// generateController 创建controller类
func generateController() error {
	moduleName := strings.ToLower(modName)
	controllerName := strings.ToLower(ctlName)
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
	path := "app/controller/" + modName + "/" + controllerName + ".go"
	if flag := util.IsPathExist(path); !flag {
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", controllerName+".go create failed"))
			return nil
		}
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", controllerName+".go create success"))
		return nil
	}
	fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", controllerName+".go already existed"))
	return nil
}
