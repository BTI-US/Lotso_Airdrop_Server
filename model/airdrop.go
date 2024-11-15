package model

import (
	mingfugorm "github.com/Fu-XDU/mingfu_go_common/gorm"
	"gorm.io/gorm"
	"time"
)

type Airdrop struct {
	gorm.Model        `json:"-"`
	Address           Address                `gorm:"foreignKey:AddressID" json:"-"`
	AddressID         uint                   `gorm:"unique;" json:"-"`
	AirdropCount      uint64                 `gorm:"type:bigint unsigned;default:0;" json:"airdrop_count"` // 未乘以10^18的值
	ContractAddress   string                 `gorm:"index:;type:char(40);not null" json:"contract_address"`
	Proofs            mingfugorm.StringArray `json:"proofs"`
	ScheduledDelivery time.Time              `json:"scheduled_delivery"`
}
