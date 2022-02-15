package lib

import (
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/schema"
	"time"
)

type DatabaseConfig struct {
	Host         string
	Port         string
	User         string
	Pass         string
	DbName       string
	Prefix       string
	MaxIdleConns int
	MaxOpenConns int
	MaxLifeTime  int // 分钟
}

// NewMysql 数据库连接
func NewMysql(config DatabaseConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", config.User, config.Pass, config.Host, config.Port, config.DbName)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		NamingStrategy: schema.NamingStrategy{
			TablePrefix:   config.Prefix,
			SingularTable: true, // 是否设置单数表名，设置为 是
		},
	})
	if err != nil {
		return nil, fmt.Errorf("Unable to connect to the database, please check the MySQL configuration information first,the error details are:" + err.Error())
	}
	// GORM 使用 database/sql 维护连接池
	sqlDB, _ := db.DB()
	// SetMaxIdleConns 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxIdleConns(config.MaxIdleConns)
	// SetMaxOpenConns 设置打开数据库连接的最大数量。
	sqlDB.SetMaxOpenConns(config.MaxOpenConns)
	// SetConnMaxLifetime 设置了连接可复用的最大时间。
	sqlDB.SetConnMaxLifetime(time.Duration(config.MaxLifeTime))
	return db, nil
}
