package command

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	dmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"mqenergy-go/global"
	"strconv"
	"time"
)

// GenerateMigrate 执行migrate
func GenerateMigrate(step string) error {
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
			return err
		}
	} else {
		n, _ := strconv.Atoi(step)
		if err := m.Steps(n); err != nil {
			return err
		}
	}
	return nil
}
