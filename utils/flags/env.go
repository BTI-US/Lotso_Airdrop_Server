package flags

import (
	"Lotso_Airdrop_Server/utils"
	"github.com/ethereum/go-ethereum/common"
)

var (
	SslCertPath string
	SslKeyPath  string

	ApiMode         int
	ApiUrl          string
	ContractAddress string
	CutoffBlock     string
	PrivateKey      string
	ChainID         int64

	MysqlHost   string
	MysqlUser   string
	MysqlPasswd string
	MysqlDB     string
	MysqlPort   int

	Decimals            uint64
	BuyerRewardLimit    uint64
	NotBuyerRewardLimit uint64
	PairAddress         string

	Debug bool

	Contract common.Address
)

func IsValidCutoffBlock(cutoffBlock string) bool {
	switch cutoffBlock {
	case "latest", "earliest", "pending", "safe", "finalized":
		return true
	default:
		return utils.IsHex(cutoffBlock)
	}
}
