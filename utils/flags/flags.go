package flags

import (
	"Lotso_Airdrop_Server/utils"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"os"
	"strconv"
)

var (
	PortFlag = cli.StringFlag{
		Name:    "port",
		Aliases: []string{"p"},
		Usage:   "Server port",
		Value:   "1423",
		EnvVars: []string{"SERVER_PORT"},
		Action: func(ctx *cli.Context, portStr string) (err error) {
			port, err := strconv.Atoi(portStr)
			if err != nil || port <= 0 || port >= 1<<16 {
				err = fmt.Errorf("flag port value %v out of range[0-65535].", port)
			}
			return
		},
	}

	ApiModeFlag = cli.IntFlag{
		Name:        "ApiMode, a",
		Usage:       "ApiMode, must be 1 or 0",
		Value:       1,
		EnvVars:     []string{"API_MODE"},
		Destination: &ApiMode,
	}

	SslCertFlag = cli.StringFlag{
		Name:    "ssl_cert, c",
		Aliases: []string{"c"},
		Usage:   "SSL cert path",
		Value:   "",
		EnvVars: []string{"SSL_CERT"},
		Action: func(ctx *cli.Context, path string) (err error) {
			if len(path) != 0 && !FileExists(path) {
				err = fmt.Errorf("SSL cert %v is not exist.", path)
			}
			return
		},
	}

	SslKeyFlag = cli.StringFlag{
		Name:    "ssl_key",
		Aliases: []string{"k"},
		Usage:   "SSL key path",
		Value:   "",
		EnvVars: []string{"SSL_KEY"},
		Action: func(ctx *cli.Context, path string) (err error) {
			if len(path) != 0 && !FileExists(path) {
				err = fmt.Errorf("SSL key %v is not exist.", path)
			}
			return
		},
	}

	DebugFlag = cli.BoolFlag{
		Name:        "debug, d",
		Usage:       "Enable debug API",
		EnvVars:     []string{"DEBUG"},
		Destination: &Debug,
	}

	ApiUrlFlag = cli.StringFlag{
		Name:        "api",
		Usage:       "Api url which is used to connect to the RPC interface of the chain",
		EnvVars:     []string{"API_URL"},
		Destination: &ApiUrl,
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:        "private, pri",
		Usage:       "Account private key used to issue airdrops regularly, hexadecimal string type",
		EnvVars:     []string{"PRIVATE_KEY"},
		Destination: &PrivateKey,
	}
	ChainIDFlag = cli.Int64Flag{
		Name:        "chain_id",
		Usage:       "The chain ID which server will connect to",
		EnvVars:     []string{"CHAIN_ID"},
		Destination: &ChainID,
	}

	CutoffBlockFlag = cli.StringFlag{
		Name:        "CutoffBlock",
		Usage:       "When counting the number of transactions, the number of transactions before this block will be counted",
		EnvVars:     []string{"CUTOFF_BLOCK"},
		Value:       "latest",
		Destination: &CutoffBlock,
	}

	ContractAddressFlag = cli.StringFlag{
		Name:        "contract",
		Usage:       "The airdrop contract address, hexadecimal string type",
		EnvVars:     []string{"CONTRACT_ADDRESS"},
		Destination: &ContractAddress,
	}

	MysqlHostFlag = cli.StringFlag{
		Name:        "MysqlHost",
		Usage:       "Mysql host",
		Value:       "127.0.0.1",
		EnvVars:     []string{"MYSQL_HOST"},
		Destination: &MysqlHost,
	}
	MysqlPortFlag = cli.IntFlag{
		Name:        "MysqlPort",
		Usage:       "Mysql port",
		Value:       3306,
		EnvVars:     []string{"MYSQL_PORT"},
		Destination: &MysqlPort,
	}
	MysqlUserFlag = cli.StringFlag{
		Name:        "MysqlUser",
		Usage:       "Mysql user",
		Value:       "root",
		EnvVars:     []string{"MYSQL_USER"},
		Destination: &MysqlUser,
	}
	MysqlPasswdFlag = cli.StringFlag{
		Name:        "MysqlPasswd",
		Usage:       "Mysql password",
		Value:       "123456",
		EnvVars:     []string{"MYSQL_PASSWD"},
		Destination: &MysqlPasswd,
	}
	MysqlDBFlag = cli.StringFlag{
		Name:        "MysqlDB",
		Usage:       "Mysql database",
		Value:       "lotso",
		EnvVars:     []string{"MYSQL_DB"},
		Destination: &MysqlDB,
	}

	TrustedProxiesFlag = cli.StringSliceFlag{
		Name:    "trusted_proxies",
		Usage:   "Trusted proxies, example: --trusted_proxies 127.0.0.1,172.0.0.1",
		EnvVars: []string{"TRUSTED_PROXIES"},
		Value:   cli.NewStringSlice(),
		Action: func(ctx *cli.Context, trustedProxies []string) (err error) {
			if len(trustedProxies) != 0 {
				log.Printf("Trust proxies: %v.", trustedProxies)
			}
			return
		},
	}

	PairAddressFlag = cli.StringFlag{
		Name:        "pairAddr",
		Usage:       "The uniswapPair address used to count buyers",
		EnvVars:     []string{"PAIR_ADDRESS"},
		Destination: &PairAddress,
	}

	DecimalsFlag = cli.Uint64Flag{
		Name:        "decimals",
		Usage:       "Decimals of token",
		Value:       18,
		EnvVars:     []string{"DECIMALS"},
		Destination: &Decimals,
	}

	BuyerRewardLimitFlag = cli.Uint64Flag{
		Name:        "buyerRewardLimit",
		Usage:       "Buyer reward limit",
		Value:       10000000,
		EnvVars:     []string{"BUYER_REWARD_LIMIT"},
		Destination: &BuyerRewardLimit,
	}
	NotBuyerRewardLimitFlag = cli.Uint64Flag{
		Name:        "notBuyerRewardLimit",
		Usage:       "NotBuyer reward limit",
		Value:       2000000,
		EnvVars:     []string{"NOT_BUYER_REWARD_LIMIT"},
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

func FileExists(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}
