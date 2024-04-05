package model

import (
	"math/big"
	"time"
)

type TransactionCount struct {
	ID                int       `json:"-"`
	Address           string    `gorm:"index;type:char(40);" json:"address"`
	TransactionCount  uint64    `json:"transaction_count"`
	CreatedAt         time.Time `json:"-"`
	AirdropCount      *big.Int  `gorm:"-" json:"airdrop_count"`
	HasAirdropped     bool      `json:"has_airdropped"`
	ScheduledDelivery time.Time `json:"scheduled_delivery"`
}
