package model

import (
	"github.com/ethereum/go-ethereum/common"
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

func NewTransactionCount(ID int, address string, transactionCount uint64, createdAt time.Time, airdropCount *big.Int, hasAirdropped bool, scheduledDelivery time.Time) *TransactionCount {
	return &TransactionCount{ID: ID, Address: address, TransactionCount: transactionCount, CreatedAt: createdAt, AirdropCount: airdropCount, HasAirdropped: hasAirdropped, ScheduledDelivery: scheduledDelivery}
}

func NewAirdrop(address *common.Address, airdropCount *big.Int) *TransactionCount {
	return &TransactionCount{Address: address.Hex(), AirdropCount: airdropCount}
}
