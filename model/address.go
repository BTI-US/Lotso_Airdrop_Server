package model

import (
	"gorm.io/gorm"
)

type Address struct {
	gorm.Model       `json:"-"`
	Address          string `gorm:"index:;type:char(40);unique;not null" json:"address"`
	TransactionCount uint64 `gorm:"type:bigint unsigned;default:0;" json:"transaction_count"`
}
