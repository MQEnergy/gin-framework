package bootstrap

import (
	"fmt"
	"github.com/MQEnergy/gin-framework/config"
	"github.com/MQEnergy/gin-framework/global"
	"github.com/MQEnergy/gin-framework/pkg/lib"
	"github.com/MQEnergy/gin-framework/pkg/util"
	"github.com/sirupsen/logrus"
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
		MysqlService:  bootMysql,
		RedisService:  bootRedis,
		LoggerService: bootLogger,
	}
)

// BootService 加载服务
func BootService(services ...string) {
	// 初始化配置
	global.Cfg = config.InitConfig()

	serviceMap[LoggerService] = bootLogger
	if global.Logger != nil {
		global.Logger.Infof("服务列表已加载完成")
	}
	if len(services) == 0 {
		services = serviceMap.keys()
	}
	BootedService = make([]string, 0)
	for k, val := range serviceMap {
		if util.InAnySlice[string](services, k) {
			if err := val(); err != nil {
				panic("程序服务启动失败:" + err.Error())
			}
			BootedService = append(BootedService, k)
		}
	}
}

// bootLogger 将配置载入日志服务
func bootLogger() error {
	if global.Logger != nil {
		return nil
	}
	if global.Logger, err = lib.NewLogger(global.Cfg.Log.DirPath, global.Cfg.Log.FileName, global.Cfg.Log.Debug); err == nil {
		logrus.Printf("程序载入Logger服务成功 [ 日志名：%s 日志路径：%s ]", global.Cfg.Log.FileName, global.Cfg.Log.DirPath)
	}
	return err
}

// bootMysql 装配数据库连接
func bootMysql() error {
	if global.DB != nil {
		return nil
	}
	dbConfig := lib.DatabaseConfig{
		Host:         global.Cfg.Mysql[0].Host,
		Port:         global.Cfg.Mysql[0].Port,
		User:         global.Cfg.Mysql[0].User,
		Pass:         global.Cfg.Mysql[0].Password,
		DbName:       global.Cfg.Mysql[0].DbName,
		Prefix:       global.Cfg.Mysql[0].Prefix,
		MaxIdleConns: global.Cfg.Mysql[0].MaxIdleConns,
		MaxOpenConns: global.Cfg.Mysql[0].MaxOpenConns,
		MaxLifeTime:  global.Cfg.Mysql[0].MaxLifeTime,
	}
	global.DB, err = lib.NewMysql(dbConfig)
	if err == nil {
		logrus.Printf("程序载入MySQL服务成功")
	}
	return err
}

// bootRedis 装配redis服务
func bootRedis() error {
	redisConfig := lib.RedisConfig{
		Addr:     fmt.Sprintf("%s:%s", global.Cfg.Redis.Host, global.Cfg.Redis.Port),
		Password: global.Cfg.Redis.Password,
		DbNum:    global.Cfg.Redis.DbNum,
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
