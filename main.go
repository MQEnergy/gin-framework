package main

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/urfave/cli/v2"
	"mqenergy-go/bootstrap"
	"mqenergy-go/command"
	"mqenergy-go/config"
	"mqenergy-go/pkg/validator"
	"mqenergy-go/router"
	"os"
	"runtime"
)

var (
	// AppName 当前应用名称
	AppName  = "gin-framework-template"
	AppUsage = "使用gin框架作为基础开发库，封装一套适用于面向api编程的快速开发框架"
	// 作者
	AuthorName  = "chenxi"
	AuthorEmail = "bbxycx.18@163.com"
	//	AppPort 程序启动端口
	AppPort string
	//	BuildVersion 编译的app版本
	BuildVersion string
	//	BuildAt 编译时间
	BuildAt string
)

// Stack 程序运行前的处理
func Stack() *cli.App {
	buildInfo := fmt.Sprintf("%s-%s-%s-%s-%s", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildAt, gtime.Now())

	return &cli.App{
		Name:    AppName,
		Version: buildInfo,
		Usage:   AppUsage,
		Authors: []*cli.Author{
			{
				Name:  AuthorName,
				Email: AuthorEmail,
			},
		},
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | test | prod]",
				Destination: &config.ConfEnv,
			},
			&cli.StringFlag{
				Name:        "port",
				Value:       "9527",
				Usage:       "请选择启动端口",
				Destination: &AppPort,
			},
		},
		Action: func(context *cli.Context) error {
			//	初始化配置文件信息
			config.InitConfig()
			//	程序启动时需要加载的服务
			bootstrap.BootService()
			//	引入验证翻译器
			validator.NewValidate()
			//	注册路由 启动程序
			return router.Register().Run(":" + AppPort)
		},
		Commands: []*cli.Command{
			{
				Name:  "migrate",
				Usage: "Create migration command",
				Action: func(ctx *cli.Context) error {
					command.GenerateMigrate()
					return nil
				},
			},
			{
				Name:  "account",
				Usage: "Create a new admin account",
				Action: func(ctx *cli.Context) error {
					command.GenerateAdmin()
					return nil
				},
			},
			{
				Name:  "model",
				Usage: "Create a new model class",
				Action: func(ctx *cli.Context) error {
					command.GenerateModel()
					return nil
				},
			},
			{
				Name:  "controller",
				Usage: "Create a new controller class",
				Action: func(ctx *cli.Context) error {
					command.GenerateController()
					return nil
				},
			},
			{
				Name:  "service",
				Usage: "Create a new service class",
				Action: func(ctx *cli.Context) error {
					command.GenerateService()
					return nil
				},
			},
		},
	}
}

func main() {
	if err := Stack().Run(os.Args); err != nil {
		panic(err)
	}
}
