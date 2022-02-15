package models

type User struct {
	Model
	Nickname string `gorm:"column:nickname;type:varchar(80);NOT NULL;comment:昵称" json:"nickname"`
	Phone    string `gorm:"column:phone;type:varchar(11);NOT NULL;comment:手机号" json:"phone"`
	Password string `gorm:"column:password;type:varchar(200);NOT NULL;comment:密码" json:"password"`
	Status   *int8  `gorm:"column:status;type:tinyint(1);NOT NULL;comment:状态1：正常 0：不正常" json:"status"`
}
