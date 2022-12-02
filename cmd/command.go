package cmd

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/bootstrap"
	"github.com/MQEnergy/gin-framework/pkg/util"
	"github.com/urfave/cli/v2"
	"os"
	"strings"
)

var (
	name string
	dir  string
)

// CommandCmd 创建command工具
func CommandCmd() *cli.Command {
	return &cli.Command{
		Name:  "command",
		Usage: "Create a new command",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:        "name",
				Aliases:     []string{"n"},
				Value:       "",
				Usage:       "请输入命令工具名称 如：command",
				Destination: &name,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "dir",
				Aliases:     []string{"d"},
				Value:       "",
				Usage:       "请输入命令工具目录 如：test",
				Destination: &dir,
				Required:    false,
			},
		},
		Action: func(c *cli.Context) error {
			bootstrap.BootService(bootstrap.LoggerService)
			return generateCommand()
		},
	}
}

// generateCommand 生成command命令工具
func generateCommand() error {
	cmdName := strings.ToLower(name)
	cmdDir := strings.ToLower(dir)
	firstUpperCtlName := strings.ToUpper(cmdName[:1]) + cmdName[1:]
	projectModuleName := util.GetProjectModuleName()
	path := "cmd/"
	packageName := "cmd"
	if cmdDir != "" {
		path += cmdDir + "/"
		if err := util.MakeMultiDir(path); err != nil {
			return err
		}
		packageName = cmdDir
	}
	path += cmdName + ".go"
	content := fmt.Sprintf(`package %s

import (
	"%s/bootstrap"
	"%s/config"
	"github.com/urfave/cli/v2"
)

var (
	%sParams string
)

// %sCmd command工具创建命令
func %sCmd() *cli.Command {
	return &cli.Command{
        Name:  "%s",
        Usage: "Create a new %s",
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
				Name:        "%sParams",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "参数值",
				Destination: &%sParams,
				Required:    true,
			},
		},
		Action: func(c *cli.Context) error {
			bootstrap.BootService() // 可按需引用服务 bootstrap.LoggerService, bootstrap.MysqlService, bootstrap.RedisService
			return generate%s()
		},
	}
}

// generate%s generate command
func generate%s() error {
	return nil
}
`,
		packageName, projectModuleName, projectModuleName, cmdName,
		firstUpperCtlName, firstUpperCtlName, cmdName, firstUpperCtlName,
		cmdName, cmdName, firstUpperCtlName, firstUpperCtlName, firstUpperCtlName)

	if flag := util.IsPathExist(path); !flag {
		if err := os.WriteFile(path, []byte(content), 0644); err != nil {
			fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", cmdName+".go create failed"))
			return nil
		}
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", cmdName+".go create success"))
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "1、需要在main.go的Commands中引用如下："+packageName+"."+firstUpperCtlName+"Cmd()"))
		fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "2、查看帮助：go run main.go "+cmdName+" --help"))
		return nil
	}
	fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", cmdName+".go already existed"))
	return nil
}
