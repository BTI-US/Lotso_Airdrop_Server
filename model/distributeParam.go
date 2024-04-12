package model

import "math/big"

type DistributeParam struct {
	Address string  `json:"address"`
	Amount  big.Int `json:"amount"`
}

type ClaimParam struct {
	PrivateKey string `json:"private_key"`
}
