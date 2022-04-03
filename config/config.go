package config

import (
	"fmt"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"os"
)

var (
	Conf    *Yaml
	ConfEnv string
)

type Yaml struct {
	Server `yaml:"server"`
	Log    `yaml:"log"`
	Mysql  `yaml:"mysql"`
	Redis  `yaml:"redis"`
	Oss    `yaml:"oss"`
}

type Server struct {
	Mode            string `yaml:"mode"`
	DefaultPageSize int    `yaml:"defaultPageSize"`
	MaxPageSize     int    `yaml:"maxPageSize"`
	FileUploadPath  string `yaml:"fileUploadPath"`
	TokenExpire     int64  `yaml:"tokenExpire"`
	TokenKey        string `yaml:"tokenKey"`
	TokenIssuer     string `yaml:"tokenIssuer"`
	JwtSecret       string `yaml:"jwtSecret"`
}

type Log struct {
	Debug    string `yaml:"debug"`
	FileName string `yaml:"fileName"`
	DirPath  string `yaml:"dirPath"`
}

type Mysql struct {
	Host         string `yaml:"host"`
	Port         string `yaml:"port"`
	User         string `yaml:"user"`
	Pass         string `yaml:"password"`
	DbName       string `yaml:"dbname"`
	Prefix       string `yaml:"prefix"`
	MaxIdleConns int    `yaml:"maxIdleConns"` // 设置空闲连接池中连接的最大数量
	MaxOpenConns int    `yaml:"maxOpenConns"` // 设置打开数据库连接的最大数量
	MaxLifeTime  int    `yaml:"maxLifeTime"`  // 设置了连接可复用的最大时间（分钟）
}

type Redis struct {
	Host        string `yaml:"host"`
	Port        string `yaml:"port"`
	Password    string `yaml:"password"`
	DbNum       int    `yaml:"dbNum"`
	LoginPrefix string `yaml:"loginPrefix"`
}

type Oss struct {
	EndPoint        string `yaml:"endPoint"`
	AccessKeyID     string `yaml:"accessKeyID"`
	AccessKeySecret string `yaml:"accessKeySecret"`
	BucketName      string `yaml:"bucketName"`
}

func InitConfig() {
	var configFile = fmt.Sprintf("config.%s.yaml", ConfEnv)
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(fmt.Errorf("读取配置文件失败：%s", err))
	}
	//	根据当前环境的值来替换配置文件中的环境变量（配合docker）
	yamlConf = []byte(os.ExpandEnv(string(yamlConf)))
	c := &Yaml{}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		panic(fmt.Errorf("解析配置文件失败：%s", err))
	}
	Conf = c
}
