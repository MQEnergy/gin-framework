package config

import (
	"github.com/jinzhu/configor"
)

var (
	ConfEnv string = "dev"
)

var Conf = struct {
	Server struct {
		Mode            string `yaml:"mode" default:"debug"`
		DefaultPageSize int    `yaml:"default_page_size" default:"10"`
		MaxPageSize     int    `yaml:"max_page_size" default:"500"`
		FileUploadPath  string `yaml:"file_upload_path" default:""`
	}
	Jwt struct {
		TokenExpire int64  `yaml:"token_expire" default:"864000"`
		TokenKey    string `yaml:"token_key" default:"Authorization"`
		TokenIssuer string `yaml:"token_issuer" default:"gin-framework"`
		Secret      string `yaml:"secret"`
	}
	Log struct {
		Debug    bool   `yaml:"debug" default:"true"`
		FileName string `yaml:"file_name" default:"gin-framework"`
		DirPath  string `yaml:"dir_path" default:"runtime/logs"`
	}
	Mysql []struct {
		Host         string `yaml:"host" default:"127.0.0.1"`
		Port         string `yaml:"port" default:"3306"`
		User         string `yaml:"user" default:"root"`
		Password     string `yaml:"password" default:"123456"`
		DbName       string `yaml:"dbname"`
		Prefix       string `yaml:"prefix" default:""`
		MaxIdleConns int    `yaml:"max_idle_conns" default:"10"`
		MaxOpenConns int    `yaml:"max_open_conns" default:"100"`
		MaxLifeTime  int    `yaml:"max_life_time" default:"60"`
	}
	Redis struct {
		Host        string `yaml:"host" default:"127.0.0.1"`
		Port        string `yaml:"port" default:"6379"`
		Password    string `yaml:"password"`
		DbNum       int    `yaml:"db_num" default:"0"`
		LoginPrefix string `yaml:"login_prefix" default:"mqenergy_login_auth_"`
	}
	Amqp struct {
		Host     string `yaml:"host" default:"127.0.0.1"`
		Port     string `yaml:"port" default:"5672"`
		User     string `yaml:"user" default:"guest"`
		Password string `yaml:"password" default:""`
		Vhost    string `yaml:"vhost" default:""`
	}
	Oss struct {
		EndPoint        string `yaml:"end_point" default:"https://oss-cn-shanghai.aliyuncs.com"`
		AccessKeyId     string `yaml:"access_key_id"`
		AccessKeySecret string `yaml:"access_key_secret"`
		BucketName      string `yaml:"bucket_name"`
	}
}{}

func InitConfig() {
	configor.Load(&Conf, "config."+ConfEnv+".yaml")
}
