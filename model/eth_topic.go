package model

type TransferTopic struct {
	ID              int    `gorm:"type:bigint;not null" json:"id"`
	BlockHash       string `gorm:"type:char(64);not null" json:"block_hash"`
	BlockNumber     uint64 `gorm:"type:bigint unsigned;not null" json:"block_number"`
	TransactionHash string `gorm:"type:char(64);not null" json:"transaction_hash"`
	From            string `gorm:"index;type:char(40);not null" json:"from"`
	To              string `gorm:"index;type:char(40);not null" json:"to"`
	Data            string `gorm:"type:char(64);not null" json:"data"`
}
