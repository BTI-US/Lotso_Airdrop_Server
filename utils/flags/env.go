package flags

import (
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
)

var (
	ApiUrl      string
	CutoffBlock string
	ChainID     int64

	Decimals    uint
	PairAddress string

	Debug bool
)

var (
	PrivateKey   *ecdsa.PrivateKey
	TokenAddress common.Address
)
