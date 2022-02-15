package entities

import "lyky-go/models"

type BaseUser models.GinUser

type User struct {
	BaseUser
	RoleIds []uint64 `gorm:"-" json:"role_ids"`
}

// UserLoginRequest 用户登录参数
type UserLoginRequest struct {
	Phone    string `json:"phone"`
	Password string `json:"password"`
}
