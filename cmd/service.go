package cmd

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"io/ioutil"
	"mqenergy-go/bootstrap"
	"mqenergy-go/config"
	"mqenergy-go/pkg/util"
	"strings"
)

var (
	serviceName string
	moduleName  string
)

// ServiceCmd 数据库模型创建命令
func ServiceCmd() *cli.Command {
	return &cli.Command{
		Name:  "service",
		Usage: "Create a new service class",
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
				Name:        "service",
				Aliases:     []string{"s"},
				Value:       "",
				Usage:       "请输入服务类名称 如：admin",
				Destination: &serviceName,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "module",
				Aliases:     []string{"m"},
				Value:       "backend",
				Usage:       "请输入模块名称 如：backend",
				Destination: &moduleName,
				Required:    false,
			},
		},
		Action: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.LoggerService)
			return generateService()
		},
	}
}

// generateService 创建service类
func generateService() error {
	moduleName := strings.ToLower(moduleName)
	serviceName := strings.ToLower(serviceName)
	firstUpperCtlName := strings.ToUpper(serviceName[:1]) + serviceName[1:]
	content := `package ` + moduleName + `

type ` + firstUpperCtlName + `Service struct{}

var ` + firstUpperCtlName + ` = &` + firstUpperCtlName + `Service{}

// Index 获取列表
func (s *` + firstUpperCtlName + `Service) Index(requestParams interface{}) (interface{}, error) {
	// Todo list 

	return nil, nil
}
`
	path := "app/service/" + moduleName + "/" + serviceName + ".go"
	if flag := util.IsPathExist(path); !flag {
		if err := ioutil.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", serviceName+".go create failed"))
			return nil
		}
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", serviceName+".go create success"))
		return nil
	}
	fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", serviceName+".go already existed"))
	return nil
}
