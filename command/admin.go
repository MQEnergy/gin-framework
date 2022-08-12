package command

import (
	"fmt"
	"mqenergy-go/bootstrap"
	"mqenergy-go/config"
	"mqenergy-go/global"
	"mqenergy-go/models"
	"mqenergy-go/pkg/util"
	"os"
	"time"
)

type Admin models.GinAdmin

// GenerateAdmin 生成admin信息工具
func GenerateAdmin() {
	var env = "dev"
	args := os.Args
	if len(args) < 4 {
		fmt.Println("参数缺失：需要两个参数 {account} {password}")
		return
	}
	if len(args) > 4 {
		env = args[4]
	}
	account := args[2]
	password := args[3]
	config.ConfEnv = env
	bootstrap.BootService("Mysql")

	salt := util.GenerateUuid(32)
	user := Admin{
		Uuid:         util.GenerateBaseSnowId(32),
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
		Status:       1,
		CreatedAt:    uint64(time.Now().Unix()),
		UpdatedAt:    uint64(time.Now().Unix()),
	}
	err := global.DB.Create(&user).Error
	if err != nil {
		panic(err)
	}
	fmt.Println("账号：" + account + "; 密码：" + password + " 生成成功")
	return
}
