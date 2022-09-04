package models

var GinUserTbName = "gin_user"

// GinUser 用户表
type GinUser struct {
	Id           uint64 `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;NOT NULL;" json:"id"`
	Uuid         string `gorm:"column:uuid;type:varchar(32);NOT NULL;comment:唯一id号" json:"uuid"`                           // 唯一id号
	UserName     string `gorm:"column:user_name;type:varchar(64);NOT NULL;comment:用户名" json:"user_name"`                   // 用户名
	NickName     string `gorm:"column:nick_name;type:varchar(64);NOT NULL;comment:昵称" json:"nick_name"`                    // 昵称
	RealName     string `gorm:"column:real_name;type:varchar(64);NOT NULL;comment:真实姓名" json:"real_name"`                  // 真实姓名
	Phone        string `gorm:"column:phone;type:varchar(16);NOT NULL;comment:手机号" json:"phone"`                           // 手机号
	Avatar       string `gorm:"column:avatar;type:varchar(128);NOT NULL;comment:头像" json:"avatar"`                         // 头像
	Password     string `gorm:"column:password;type:varchar(64);NOT NULL;comment:密码" json:"password"`                      // 密码
	Salt         string `gorm:"column:salt;type:varchar(32);NOT NULL;comment:密码" json:"salt"`                              // 密码
	Status       uint8  `gorm:"column:status;type:tinyint unsigned;default:1;NOT NULL;comment:状态 1：正常 2：禁用" json:"status"` // 状态 1：正常 2：禁用
	RegisterTime string `gorm:"column:register_time;type:varchar(16);NOT NULL;comment:注册时间" json:"register_time"`          // 注册时间
	RegisterIp   string `gorm:"column:register_ip;type:varchar(32);NOT NULL;comment:注册ip" json:"register_ip"`              // 注册ip
	LoginTime    string `gorm:"column:login_time;type:varchar(16);NOT NULL;comment:登录时间" json:"login_time"`                // 登录时间
	LoginIp      string `gorm:"column:login_ip;type:varchar(32);NOT NULL;comment:登录ip" json:"login_ip"`                    // 登录ip
	CreatedAt    uint64 `gorm:"column:created_at;type:bigint unsigned;NULL;" json:"created_at"`
	UpdatedAt    uint64 `gorm:"column:updated_at;type:bigint unsigned;NULL;" json:"updated_at"`
}
