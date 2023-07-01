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
	roleId   string
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
			&cli.StringFlag{
				Name:        "role",
				Aliases:     []string{"r"},
				Value:       "",
				Usage:       "请输入角色ID 如：1,2,3.. 或连续的1,2,3",
				Destination: &roleId,
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
	var (
		adminInfo models.GinAdmin
		salt      = util.GenerateUuid(32)
		uuid      = util.GenerateBaseSnowId(0, nil)
		pass      = util.GeneratePasswordHash(password, salt)
		timeNow   = uint64(time.Now().Unix())
	)
	if err := global.DB.First(&adminInfo, "account = ?", account).Error; err == nil {
		fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", "account: "+account+" is already existed"))
		return err
	}
	localIp, err := util.GetLocalIp()
	if err != nil {
		return err
	}
	adminInfo = models.GinAdmin{
		Uuid:         uuid,
		Account:      account,
		Password:     pass,
		Phone:        "",
		Avatar:       "",
		Salt:         salt,
		RealName:     account,
		RegisterTime: timeNow,
		RegisterIp:   localIp,
		LoginTime:    0,
		LoginIp:      "",
		RoleIds:      roleId,
		Status:       1,
		CreatedAt:    timeNow,
		UpdatedAt:    timeNow,
	}
	if err := global.DB.Create(&adminInfo).Error; err != nil {
		return err
	}
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", fmt.Sprintf("账号：%s 密码：%s 角色ID：%s 生成成功", account, password, roleId)))
	return nil
}
