package cmd

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/bootstrap"
	"github.com/MQEnergy/gin-framework/config"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/models"
	"github.com/MQEnergy/gin-framework/pkg/util"
	"github.com/urfave/cli/v2"
	"time"
)

var (
	account  string
	password string
)

// AccountCmd 管理者账号创建命令
func AccountCmd() *cli.Command {
	return &cli.Command{
		Name:  "account",
		Usage: "Create a new manager account",
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
				Name:        "account",
				Aliases:     []string{"a"},
				Value:       "",
				Usage:       "请输入账号名称 如：admin",
				Destination: &account,
				Required:    true,
			},
			&cli.StringFlag{
				Name:        "password",
				Aliases:     []string{"p"},
				Value:       "",
				Usage:       "请输入账号密码 如：admin888",
				Destination: &password,
				Required:    true,
			},
		},
		Action: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.MysqlService)
			return generateAdmin()
		},
	}
}

// generateAdmin 生成admin信息工具
func generateAdmin() error {
	salt := util.GenerateUuid(32)
	admin := &models.GinAdmin{
		Uuid:         util.GenerateBaseSnowId(32, nil),
		Account:      account,
		Password:     util.GeneratePasswordHash(password, salt),
		Phone:        "12345678901",
		Avatar:       "",
		Salt:         salt,
		RealName:     account,
		RegisterTime: uint64(time.Now().Unix()),
		RegisterIp:   "127.0.0.1",
		LoginTime:    uint64(time.Now().Unix()),
		LoginIp:      "127.0.0.1",
		RoleIds:      "1",
		Status:       1,
		CreatedAt:    uint64(time.Now().Unix()),
		UpdatedAt:    uint64(time.Now().Unix()),
	}
	if err := global.DB.Create(admin).Error; err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "账号："+account+"; 密码："+password+" 生成成功"))
	return nil
}
