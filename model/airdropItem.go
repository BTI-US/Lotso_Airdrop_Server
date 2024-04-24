package model

import (
	"github.com/ethereum/go-ethereum/common"
	"time"
)

type AirdropItem struct {
	ID                uint      `json:"-"`
	Address           string    `gorm:"index:;type:char(40);unique;not null" json:"address"`
	TransactionCount  uint64    `json:"transaction_count"`
	CreatedAt         time.Time `json:"-"`
	AirdropCount      uint64    `gorm:"type:int;" json:"airdrop_count"`
	HasAirdropped     bool      `json:"has_airdropped"`
	ScheduledDelivery time.Time `json:"scheduled_delivery"`
}

func NewAirdropItem(ID uint, address string, transactionCount uint64, createdAt time.Time, airdropCount uint64, hasAirdropped bool, scheduledDelivery time.Time) *AirdropItem {
	return &AirdropItem{ID: ID, Address: address, TransactionCount: transactionCount, CreatedAt: createdAt, AirdropCount: airdropCount, HasAirdropped: hasAirdropped, ScheduledDelivery: scheduledDelivery}
}

func NewAirdrop(address *common.Address, airdropCount uint64) *AirdropItem {
	return &AirdropItem{Address: address.Hex(), AirdropCount: airdropCount}
}
