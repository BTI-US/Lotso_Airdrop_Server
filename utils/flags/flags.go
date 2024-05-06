package flags

import (
	"Lotso_Airdrop_Server/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gopkg.in/urfave/cli.v1"
)

var (
	PortFlag = cli.StringFlag{
		Name:   "port, p",
		Usage:  "Server port",
		Value:  "1423",
		EnvVar: "SERVER_PORT",
	}
	ApiModeFlag = cli.IntFlag{
		Name:        "ApiMode, a",
		Usage:       "ApiMode, must be 1 or 0",
		Value:       1,
		EnvVar:      "API_MODE",
		Destination: &ApiMode,
	}

	SslCertFlag = cli.StringFlag{
		Name:        "ssl_cert, c",
		Usage:       "SSL cert path",
		Value:       "",
		EnvVar:      "SSL_CERT",
		Destination: &SslCertPath,
	}

	SslKeyFlag = cli.StringFlag{
		Name:        "ssl_key, k",
		Usage:       "SSL key path",
		Value:       "",
		EnvVar:      "SSL_KEY",
		Destination: &SslKeyPath,
	}

	DebugFlag = cli.BoolFlag{
		Name:        "debug, d",
		Usage:       "Enable debug API",
		EnvVar:      "DEBUG",
		Destination: &Debug,
	}

	ApiUrlFlag = cli.StringFlag{
		Name:        "api",
		Usage:       "Api url which is used to connect to the RPC interface of the chain",
		EnvVar:      "API_URL",
		Destination: &ApiUrl,
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:        "private, pri",
		Usage:       "Account private key used to issue airdrops regularly, hexadecimal string type",
		EnvVar:      "PRIVATE_KEY",
		Destination: &PrivateKey,
	}
	ChainIDFlag = cli.Int64Flag{
		Name:        "chain_id",
		Usage:       "The chain ID which server will connect to",
		EnvVar:      "CHAIN_ID",
		Destination: &ChainID,
	}

	CutoffBlockFlag = cli.StringFlag{
		Name:        "CutoffBlock",
		Usage:       "When counting the number of transactions, the number of transactions before this block will be counted",
		EnvVar:      "CUTOFF_BLOCK",
		Value:       "latest",
		Destination: &CutoffBlock,
	}

	ContractAddressFlag = cli.StringFlag{
		Name:        "contract",
		Usage:       "The airdrop contract address, hexadecimal string type",
		EnvVar:      "CONTRACT_ADDRESS",
		Destination: &ContractAddress,
	}

	MysqlHostFlag = cli.StringFlag{
		Name:        "MysqlHost",
		Usage:       "Mysql host",
		Value:       "127.0.0.1",
		EnvVar:      "MYSQL_HOST",
		Destination: &MysqlHost,
	}
	MysqlPortFlag = cli.IntFlag{
		Name:        "MysqlPort",
		Usage:       "Mysql port",
		Value:       3306,
		EnvVar:      "MYSQL_PORT",
		Destination: &MysqlPort,
	}
	MysqlUserFlag = cli.StringFlag{
		Name:        "MysqlUser",
		Usage:       "Mysql user",
		Value:       "root",
		EnvVar:      "MYSQL_USER",
		Destination: &MysqlUser,
	}
	MysqlPasswdFlag = cli.StringFlag{
		Name:        "MysqlPasswd",
		Usage:       "Mysql password",
		Value:       "123456",
		EnvVar:      "MYSQL_PASSWD",
		Destination: &MysqlPasswd,
	}
	MysqlDBFlag = cli.StringFlag{
		Name:        "MysqlDB",
		Usage:       "Mysql database",
		Value:       "lotso",
		EnvVar:      "MYSQL_DB",
		Destination: &MysqlDB,
	}

	TrustedProxiesFlag = cli.StringSliceFlag{
		Name:   "TrustedProxies",
		Usage:  "Trusted proxies, example: --TrustedProxies 127.0.0.1,172.0.0.1",
		EnvVar: "TRUSTED_PROXIES",
	}

	PairAddressFlag = cli.StringFlag{
		Name:        "pairAddr",
		Usage:       "The uniswapPair address used to count buyers",
		EnvVar:      "PAIR_ADDRESS",
		Destination: &PairAddress,
	}

	DecimalsFlag = cli.Uint64Flag{
		Name:        "decimals",
		Usage:       "Decimals of token",
		Value:       18,
		EnvVar:      "DECIMALS",
		Destination: &Decimals,
	}

	BuyerRewardLimitFlag = cli.Uint64Flag{
		Name:        "buyerRewardLimit",
		Usage:       "Buyer reward limit",
		Value:       10000000,
		EnvVar:      "BUYER_REWARD_LIMIT",
		Destination: &BuyerRewardLimit,
	}
	NotBuyerRewardLimitFlag = cli.Uint64Flag{
		Name:        "notBuyerRewardLimit",
		Usage:       "NotBuyer reward limit",
		Value:       2000000,
		EnvVar:      "NOT_BUYER_REWARD_LIMIT",
		Destination: &NotBuyerRewardLimit,
	}
)

func ProcessFlags(ctx *cli.Context) {
	CheckRequiredFlags()
	AirdropContractAddress = common.HexToAddress(ContractAddress)
	PairAddress = utils.Remove0xPrefix(PairAddress)
	if !utils.Has0xPrefix(CutoffBlock) {
		CutoffBlock = "0x" + CutoffBlock
	}
	PrivateKey = utils.Remove0xPrefix(PrivateKey)

	TrustedProxies = ctx.StringSlice("TrustedProxies")
}

func CheckRequiredFlags() {
	SslEnabled = len(SslCertPath) != 0 && len(SslKeyPath) != 0

	if ApiMode != 0 && ApiMode != 1 {
		log.Fatalf("Wrong value for --ApiMode: %v, must be 1 or 0", ApiMode)
	}

	if len(ApiUrl) == 0 {
		log.Fatalf("--ApiUrl is Required")
	}

	if ChainID == 0 {
		log.Fatalf("--ChainID is Required")
	}

	if !IsValidCutoffBlock(CutoffBlock) {
		log.Fatalf("Wrong value for --CutoffBlock: %v, must be 'latest', 'earliest', 'pending', 'safe', 'finalized' or a block number in hexadecimal form", CutoffBlock)
	}

	if !common.IsHexAddress(ContractAddress) {
		log.Fatalf("Wrong value for --contract: %v, must be a hexadecimal address", ApiMode)
	}

	if !common.IsHexAddress(PairAddress) {
		log.Fatalf("Wrong value for --pairAddr: %v, must be a hexadecimal address", ApiMode)
	}
}

func IsValidCutoffBlock(cutoffBlock string) bool {
	switch cutoffBlock {
	case "latest", "earliest", "pending", "safe", "finalized":
		return true
	default:
		return utils.IsHex(cutoffBlock)
	}
}
