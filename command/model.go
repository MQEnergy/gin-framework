package command

import (
	"fmt"
	"lyky-go/bootstrap"
	"lyky-go/config"
	"os"
)

// GenerateModel 创建model
func GenerateModel() {
	env := "dev"
	args := os.Args
	if len(args) < 4 {
		fmt.Println("参数缺失：至少需要一个参数 {n} {env}")
		return
	}
	if len(args) >= 4 {
		env = args[3]
	}
	config.ConfEnv = env
	config.InitConfig()
	bootstrap.BootMysql()

	if args[2] == "all" {
		GenerateAllModel(config.Conf.Mysql.DbName)
	} else {
		var table Table
		table = GetSingleTable(config.Conf.Mysql.DbName, args[2])
		err := GenerateSingleModel(config.Conf.Mysql.DbName, args[2], table)
		if err != nil {
			fmt.Println(err)
		}
	}
}
