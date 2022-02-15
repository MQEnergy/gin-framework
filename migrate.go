package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/golang-migrate/migrate/v4"
	dmysql "github.com/golang-migrate/migrate/v4/database/mysql"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	"log"
	"lyky-go/config"
	"os"
	"strconv"
	"time"
)

func main() {
	var env = "dev"
	n, _ := strconv.Atoi(os.Args[1])
	if len(os.Args) >= 3 {
		env = os.Args[2]
	}
	config.ConfEnv = env
	config.InitConfig()

	cfg := mysql.Config{
		DBName:               config.Conf.Mysql.DbName,
		User:                 config.Conf.Mysql.User,
		Passwd:               config.Conf.Mysql.Pass,
		Net:                  "tcp",
		Addr:                 fmt.Sprintf("%s:%s", config.Conf.Mysql.Host, config.Conf.Mysql.Port),
		AllowNativePasswords: true,
		MultiStatements:      true,
		ParseTime:            true,
		Loc:                  time.Local,
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		panic(err.Error())
	}
	driver, err := dmysql.WithInstance(db, &dmysql.Config{
		DatabaseName: config.Conf.Mysql.DbName,
	})
	if err != nil {
		panic(err.Error())
	}
	m, _ := migrate.NewWithDatabaseInstance(
		"file://./migrations",
		"mysql",
		driver,
	)
	if err := m.Steps(n); err != nil {
		log.Fatal(err.Error())
	}
}
