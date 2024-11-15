package model

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"math/big"
)

type DistributesParam struct {
	Distributes []DistributeParam `json:"distributes"`
}

func (p DistributesParam) Validate() (err error) {
	for _, d := range p.Distributes {
		if err = d.Validate(); err != nil {
			return
		}
	}
	return
}

type DistributeParam struct {
	Address string  `json:"address"`
	Amount  big.Int `json:"amount"`
}

func (p DistributeParam) Validate() (err error) {
	if !common.IsHexAddress(p.Address) {
		err = fmt.Errorf("invalid address: %s", p.Address)
		return
	}
	return
}

type ClaimParam struct {
	PrivateKey      string   `json:"private_key"`
	AirdropCount    uint64   `json:"airdrop_count"`
	ContractAddress string   `json:"contract_address"`
	Proofs          []string `json:"proofs"`
}

type RecipientInfo struct {
	RecipientCount *big.Int `json:"recipient_count"`
	AirdropAmount  *big.Int `json:"airdrop_amount"`
}
