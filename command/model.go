package command

import (
	"fmt"
	gorm_model "github.com/MQEnergy/gorm-model"
	"mqenergy-go/bootstrap"
	"mqenergy-go/config"
	"mqenergy-go/global"
	"os"
)

// GenerateModel 创建model
func GenerateModel() {
	env := "dev"
	args := os.Args
	if len(args) < 3 {
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
		gorm_model.GenerateAllModel(global.DB, config.Conf.Mysql.DbName)
	} else {
		var table gorm_model.Table
		table = gorm_model.GetSingleTable(global.DB, config.Conf.Mysql.DbName, args[2])
		err := gorm_model.GenerateSingleModel(global.DB, args[2], table)
		if err != nil {
			fmt.Println(err)
		}
	}
}
