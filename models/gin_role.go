package models

var GinRoleTbName = "gin_role"

// GinRole 角色表
type GinRole struct {
	Id        int    `gorm:"primaryKey;autoIncrement;column:id;type:int;NOT NULL;" json:"id"`
	Name      string `gorm:"column:name;type:varchar(64);NOT NULL;comment:角色名称" json:"name"`                                    // 角色名称
	Desc      string `gorm:"column:desc;type:varchar(64);NOT NULL;comment:角色描述" json:"desc"`                                    // 角色描述
	Status    int8   `gorm:"column:status;type:tinyint(1);default:1;NOT NULL;comment:状态：1正常(默认) 0停用" json:"status"`             // 状态：1正常(默认) 0停用
	RoleType  int8   `gorm:"column:role_type;type:tinyint(1);default:1;NOT NULL;comment:角色类型 1：web角色 2：app角色" json:"role_type"` // 角色类型 1：web角色 2：app角色
	CreatedAt int    `gorm:"column:created_at;type:int;NOT NULL;" json:"created_at"`
	UpdatedAt int    `gorm:"column:updated_at;type:int;NOT NULL;" json:"updated_at"`
}
