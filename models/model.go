package models

import "github.com/MQEnergy/go-framework/pkg/util"

type Model struct {
	Id        uint64          `gorm:"primaryKey;autoIncrement;column:id;type:bigint unsigned;NOT NULL;" json:"id"`
	CreatedAt util.FormatTime `json:"created_at"`
	UpdatedAt util.FormatTime `json:"updated_at"`
}
