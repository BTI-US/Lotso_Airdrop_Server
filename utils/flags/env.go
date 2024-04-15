package flags

import (
	"Lotso_Airdrop_Server/utils"
)

var (
	SslCertPath string
	SslKeyPath  string

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

	Debug bool
)

func IsValidCutoffBlock(cutoffBlock string) bool {
	switch cutoffBlock {
	case "latest", "earliest", "pending", "safe", "finalized":
		return true
	default:
		return utils.IsHex(cutoffBlock)
	}
}
