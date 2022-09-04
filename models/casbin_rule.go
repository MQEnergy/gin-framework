package models

var CasbinRuleTbName = "casbin_rule"

type CasbinRule struct {
	Id    uint64 `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;NOT NULL;" json:"id"`
	Ptype string `gorm:"column:ptype;type:varchar(100);NULL;comment:策略类型" json:"ptype"` // 策略类型
	V0    string `gorm:"column:v0;type:varchar(100);NULL;comment:角色ID" json:"v0"`       // 角色ID
	V1    string `gorm:"column:v1;type:varchar(100);NULL;comment:api路径" json:"v1"`      // api路径
	V2    string `gorm:"column:v2;type:varchar(100);NULL;comment:api访问方法" json:"v2"`    // api访问方法
	V3    string `gorm:"column:v3;type:varchar(100);NULL;" json:"v3"`
	V4    string `gorm:"column:v4;type:varchar(100);NULL;" json:"v4"`
	V5    string `gorm:"column:v5;type:varchar(100);NULL;" json:"v5"`
	V6    string `gorm:"column:v6;type:varchar(25);NULL;" json:"v6"`
	V7    string `gorm:"column:v7;type:varchar(25);NULL;" json:"v7"`
}
