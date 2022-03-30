package bootstrap

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"mqenergy-go/config"
	"mqenergy-go/global"
	"mqenergy-go/pkg/lib"
	"mqenergy-go/pkg/util"
)

// 定义服务列表
const (
	LoggerService = `Logger`
	MysqlService  = `Mysql`
	RedisService  = `Redis`
)

type bootServiceMap map[string]func() error

// BootedService 已经加载的服务
var (
	BootedService []string
	err           error
	// serviceMap 程序启动时需要自动加载的服务
	serviceMap = bootServiceMap{
		MysqlService:  BootMysql,
		RedisService:  BootRedis,
		LoggerService: BootLogger,
	}
)

// BootService 加载服务
func BootService(services ...string) {
	serviceMap[LoggerService] = BootLogger
	if global.Logger != nil {
		global.Logger.Infof("服务列表已加载完成")
	}
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	BootedService = make([]string, 0)
	for k, val := range serviceMap {
		if util.InStringSlice[string](services, k) {
			if err := val(); err != nil {
				panic("程序服务启动失败:" + err.Error())
			}
			BootedService = append(BootedService, k)
		}
	}
}

// BootLogger 将配置载入日志服务
func BootLogger() error {
	if global.Logger != nil {
		return nil
	}
	if global.Logger, err = lib.NewLogger(config.Conf.DirPath, config.Conf.FileName); err == nil {
		logrus.Printf("程序载入Logger服务成功 模块名为：%s 日志路径为：%s", config.Conf.FileName, config.Conf.DirPath)
	}
	return err
}

// BootMysql 装配数据库连接
func BootMysql() error {
	if global.DB != nil {
		return nil
	}
	dbConfig := lib.DatabaseConfig{
		Host:         config.Conf.Mysql.Host,
		Port:         config.Conf.Mysql.Port,
		User:         config.Conf.Mysql.User,
		Pass:         config.Conf.Mysql.Pass,
		DbName:       config.Conf.Mysql.DbName,
		Prefix:       config.Conf.Mysql.Prefix,
		MaxIdleConns: config.Conf.Mysql.MaxIdleConns,
		MaxOpenConns: config.Conf.Mysql.MaxOpenConns,
		MaxLifeTime:  config.Conf.Mysql.MaxLifeTime,
	}
	global.DB, err = lib.NewMysql(dbConfig)
	if err == nil {
		logrus.Printf("程序载入MySQL服务成功")
	}
	return err
}

// BootRedis 装配redis服务
func BootRedis() error {
	redisConfig := lib.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", config.Conf.Redis.Host, config.Conf.Redis.Port),
		Password: config.Conf.Redis.Password,
		DbNum:    config.Conf.Redis.DbNum,
	}
	global.Redis, err = lib.NewRedis(redisConfig)
	if err == nil {
		logrus.Printf("程序载入Redis服务成功")
	}
	return err
}

// keys 获取BootServiceMap中所有键值
func (m bootServiceMap) keys() []string {
	keys := make([]string, 0)
	for k := range m {
		keys = append(keys, k)
	}
	return keys
}
