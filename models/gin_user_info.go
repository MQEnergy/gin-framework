package models

var GinUserInfoTbName = "gin_user_info"

type GinUserInfo struct {
	Id        uint64 `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;NOT NULL;" json:"id"`
	UserId    uint64 `gorm:"column:user_id;type:bigint unsigned;NULL;comment:用户ID" json:"user_id"`        // 用户ID
	RoleIds   string `gorm:"column:role_ids;type:varchar(64);NULL;comment:角色ID 例如：1,2,3" json:"role_ids"` // 角色ID 例如：1,2,3
	CreatedAt uint64 `gorm:"column:created_at;type:bigint unsigned;NULL;" json:"created_at"`
}
