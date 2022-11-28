package main

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/bootstrap"
	"github.com/MQEnergy/gin-framework/cmd"
	"github.com/MQEnergy/gin-framework/config"
	"github.com/MQEnergy/gin-framework/pkg/validator"
	"github.com/MQEnergy/gin-framework/router"
	"github.com/gogf/gf/v2/os/gtime"
	"github.com/urfave/cli/v2"
	"os"
	"runtime"
)

var (
	// AppName 当前应用名称
	AppName  = "gin-framework"
	AppUsage = "使用gin框架作为基础开发库，封装一套适用于面向api编程的快速开发框架"
	// Authors 作者
	Authors = []*cli.Author{
		{
			Name:  "chenxi",
			Email: "bbxycx.18@163.com",
		},
	}
	//	AppPort 程序启动端口
	AppPort string
	//	BuildVersion 编译的app版本
	BuildVersion string
	//	BuildAt 编译时间
	BuildAt string
	_UI     = `
 ██████╗ ██╗███╗   ██╗      ███████╗██████╗  █████╗ ███╗   ███╗███████╗██╗    ██╗ ██████╗ ██████╗ ██╗  ██╗
██╔════╝ ██║████╗  ██║      ██╔════╝██╔══██╗██╔══██╗████╗ ████║██╔════╝██║    ██║██╔═══██╗██╔══██╗██║ ██╔╝
██║  ███╗██║██╔██╗ ██║█████╗█████╗  ██████╔╝███████║██╔████╔██║█████╗  ██║ █╗ ██║██║   ██║██████╔╝█████╔╝ 
██║   ██║██║██║╚██╗██║╚════╝██╔══╝  ██╔══██╗██╔══██║██║╚██╔╝██║██╔══╝  ██║███╗██║██║   ██║██╔══██╗██╔═██╗ 
╚██████╔╝██║██║ ╚████║      ██║     ██║  ██║██║  ██║██║ ╚═╝ ██║███████╗╚███╔███╔╝╚██████╔╝██║  ██║██║  ██╗
 ╚═════╝ ╚═╝╚═╝  ╚═══╝      ╚═╝     ╚═╝  ╚═╝╚═╝  ╚═╝╚═╝     ╚═╝╚══════╝ ╚══╝╚══╝  ╚═════╝ ╚═╝  ╚═╝╚═╝  ╚═╝
`
)

// Stack 程序运行前的处理
func Stack() *cli.App {
	buildInfo := fmt.Sprintf("%s-%s-%s-%s-%s", runtime.GOOS, runtime.GOARCH, BuildVersion, BuildAt, gtime.Now())

	return &cli.App{
		Name:    AppName,
		Version: buildInfo,
		Usage:   AppUsage,
		Authors: Authors,
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "env",
				Aliases:     []string{"e"},
				Value:       "dev",
				Usage:       "请选择配置文件 [dev | test | prod]",
				Destination: &config.ConfEnv,
			},
			&cli.StringFlag{
				Name:        "port",
				Aliases:     []string{"p"},
				Value:       "9527",
				Usage:       "请选择启动端口",
				Destination: &AppPort,
			},
		},
		Action: func(context *cli.Context) error {
			fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", _UI))

			//	程序启动时需要加载的服务
			bootstrap.BootService()
			//	引入验证翻译器
			validator.NewValidate()
			//	注册路由 启动程序
			return router.Register().Run(":" + AppPort)
		},
		Commands: []*cli.Command{
			cmd.MigrationCmd(),  // 数据库迁移
			cmd.AccountCmd(),    // 管理账号创建
			cmd.ModelCmd(),      // 模型创建
			cmd.ControllerCmd(), // 控制器创建
			cmd.ServiceCmd(),    // 服务类创建
		},
	}
}

func main() {
	if err := Stack().Run(os.Args); err != nil {
		panic(err)
	}
}
