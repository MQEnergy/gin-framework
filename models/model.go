package models

import "lyky-go/pkg/util"

type Model struct {
	ID        uint64          `gorm:"primary_key column:id comment:主键" json:"id"`
	CreatedAt util.FormatTime `json:"created_at"`
	UpdatedAt util.FormatTime `json:"updated_at"`
}
