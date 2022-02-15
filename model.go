package main

import (
	"fmt"
	"lyky-go/bootstrap"
	"lyky-go/command"
	"lyky-go/config"
	"os"
)

func main() {
	var env = "dev"
	n := os.Args[1]
	if len(os.Args) >= 3 {
		env = os.Args[2]
	}
	config.ConfEnv = env
	config.InitConfig()
	bootstrap.BootMysql()

	if n == "all" {
		command.GenerateAllModel(config.Conf.Mysql.DbName)
	} else {
		var table command.Table
		table = command.GetSingleTable(config.Conf.Mysql.DbName, n)
		err := command.GenerateSingleModel(config.Conf.Mysql.DbName, n, table)
		if err != nil {
			fmt.Println(err)
		}
	}
}
