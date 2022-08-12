package command

import (
	"fmt"
	gorm_model "github.com/MQEnergy/gorm-model"
	"github.com/sirupsen/logrus"
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
		logrus.Error("参数错误 请输入参数：操作名称（all or 表名） 环境变量（非必填 如：dev、test、prod）")
		return
	}
	if len(args) >= 4 {
		env = args[3]
	}
	config.ConfEnv = env
	config.InitConfig()
	bootstrap.BootService("Mysql")

	if args[2] == "all" {
		gorm_model.GenerateAllModel(global.DB, global.Cfg.Mysql[0].DbName)
	} else {
		var table gorm_model.Table
		table = gorm_model.GetSingleTable(global.DB, global.Cfg.Mysql[0].DbName, args[2])
		err := gorm_model.GenerateSingleModel(global.DB, args[2], table)
		if err != nil {
			fmt.Println(err)
		}
	}
}
