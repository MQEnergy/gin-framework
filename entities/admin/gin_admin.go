package admin

import "mqenergy-go/models"

type BaseAdmin models.GinAdmin

type Admin struct {
	BaseAdmin
	RoleIds []uint64 `gorm:"-" json:"role_ids"`
}
