package cmd

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	dmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"github.com/urfave/cli/v2"
	"mqenergy-go/bootstrap"
	"mqenergy-go/config"
	"mqenergy-go/global"
	"strconv"
	"time"
)

var (
	step string
)

// MigrationCmd 数据库迁移命令
func MigrationCmd() *cli.Command {
	return &cli.Command{
		Name:  "migrate",
		Usage: "Create a migration command",
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
				Name:        "step",
				Aliases:     []string{"s"},
				Usage:       "迁移/回滚步数（正数数字代表迁移步数，负数数字代表回滚步数，all代表迁移全部）",
				Destination: &step,
				Required:    true,
			},
		},
		Action: func(ctx *cli.Context) error {
			bootstrap.BootService(bootstrap.MysqlService)
			return generateMigrate()
		},
	}
}

// generateMigrate 执行migrate
func generateMigrate() error {
	cfg := mysql.Config{
		DBName:               global.Cfg.Mysql[0].DbName,
		User:                 global.Cfg.Mysql[0].User,
		Passwd:               global.Cfg.Mysql[0].Password,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", global.Cfg.Mysql[0].Host, global.Cfg.Mysql[0].Port),
		AllowNativePasswords: true,
		MultiStatements:      true,
		ParseTime:            true,
		Loc:                  time.Local,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		return err
	}
	driver, err := dmysql.WithInstance(db, &dmysql.Config{
		DatabaseName: global.Cfg.Mysql[0].DbName,
	})
	if err != nil {
		return err
	}
	m, _ := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if step == "all" {
		if err := m.Up(); err != nil {
			fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", "migration run failed err: "+err.Error()))
			return nil
		}
	} else {
		n, _ := strconv.Atoi(step)
		if err := m.Steps(n); err != nil {
			fmt.Println(fmt.Sprintf("\x1b[31m%s\x1b[0m", "step: "+step+" migration run failed err: "+err.Error()))
			return nil
		}
	}
	fmt.Println(fmt.Sprintf("\u001B[34m%s\u001B[0m", "migration run success"))
	return nil
}
