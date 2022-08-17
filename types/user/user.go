package user

import "mqenergy-go/models"

type BaseUser models.GinUser
type GinUserInfo models.GinUserInfo

type User struct {
	BaseUser
	RoleIds []uint64 `gorm:"-" json:"role_ids"`
}

// IndexRequest 获取用户列表请求参数
type IndexRequest struct {
	Page int `form:"page" json:"page" binding:"required"`
}

// UserList joins获取关联列表
type UserList struct {
	BaseUser
	GinUserInfo `gorm:"foreignKey:user_id" json:"user_info"`
}

// GinUser preload获取关联列表
type GinUser struct {
	BaseUser
	UserInfo GinUserInfo `gorm:"foreignKey:user_id" json:"user_info"`
}

// LoginRequest 用户登录请求参数
type LoginRequest struct {
	Phone    string `form:"phone" json:"phone" binding:"required"`
	Password string `form:"password" json:"password" binding:"required"`
}
