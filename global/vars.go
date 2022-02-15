package global

import (
	"github.com/go-redis/redis/v8"
	"gorm.io/gorm"
	"lyky-go/pkg/lib"
)

type Search struct {
	Column string `json:"column"` // 字段名称 如: nickname | phone
	Action string `json:"action"` // 查询方式 如: equals 表示 = | like 表示 LIKE | in 表示 IN
	Needle string `json:"needle"` // 所需条件 如: equals时对应等于的那个值 | like时表示那个关键词 | in时表示一个用逗号分隔的字符串 "2,3,4,5,6"
}

var (
	DB     *gorm.DB      // Mysql数据库
	Logger *lib.Logger   // 日志
	Redis  *redis.Client // redis连接池
)
