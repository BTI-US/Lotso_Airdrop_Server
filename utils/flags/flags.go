package flags

import (
	"Lotso_Airdrop_Server/utils"
	"errors"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/urfave/cli/v2"
)

var (
	DebugFlag = cli.BoolFlag{
		Name:        "debug",
		Aliases:     []string{"d"},
		Usage:       "Enable debug API",
		EnvVars:     []string{"DEBUG"},
		Destination: &Debug,
	}

	ApiUrlFlag = cli.StringFlag{
		Name:        "api",
		Usage:       "Api url which is used to connect to the RPC interface of the chain",
		EnvVars:     []string{"API_URL"},
		Destination: &ApiUrl,
		Required:    true,
		Action: func(ctx *cli.Context, url string) (err error) {
			return
		},
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:     "private",
		Aliases:  []string{"priv"},
		Usage:    "Account private key used to deploy airdrop contract regularly, hexadecimal string type",
		Required: true,
		EnvVars:  []string{"PRIVATE_KEY"},
		Action: func(ctx *cli.Context, key string) (err error) {
			PrivateKey, err = crypto.HexToECDSA(utils.Remove0xPrefix(key))
			if err != nil {
				err = fmt.Errorf("invalid private key, err: %v", err)
				return
			}

			return
		},
	}

	ChainIDFlag = cli.Int64Flag{
		Name:        "chain_id",
		Usage:       "The chain ID which server will connect to",
		EnvVars:     []string{"CHAIN_ID"},
		Required:    true,
		Destination: &ChainID,
	}

	CutoffBlockFlag = cli.StringFlag{
		Name:        "cutoffBlock",
		Usage:       "When counting the number of transactions, the number of transactions before this block will be counted",
		EnvVars:     []string{"CUTOFF_BLOCK"},
		Value:       "latest",
		Destination: &CutoffBlock,
		Action: func(ctx *cli.Context, cutoffBlock string) (err error) {
			switch cutoffBlock {
			case "latest", "earliest", "pending", "safe", "finalized":
				return
			default:
				if !utils.IsHex(cutoffBlock) {
					err = errors.New("invalid cutoffBlock, must be 'latest', 'earliest', 'pending', 'safe', 'finalized' or a block number as a hexadecimal string")
					return
				}

				if !utils.Has0xPrefix(CutoffBlock) {
					CutoffBlock = "0x" + CutoffBlock
				}
			}
			return
		},
	}

	PairAddressFlag = cli.StringFlag{
		Name:        "pairAddr",
		Usage:       "The uniswapPair address used to count buyers",
		EnvVars:     []string{"PAIR_ADDRESS"},
		Required:    true,
		Destination: &PairAddress,
		Action: func(ctx *cli.Context, address string) (err error) {
			if !common.IsHexAddress(address) {
				err = fmt.Errorf("wrong value for --pairAddr: %v, must be a hexadecimal address", PairAddress)
				return
			}
			PairAddress = utils.Remove0xPrefix(PairAddress)
			return
		},
	}

	TokenAddressFlag = cli.StringFlag{
		Name:     "tokenAddr",
		Usage:    "The token address used to deploy airdrop contract",
		EnvVars:  []string{"TOKEN_ADDRESS"},
		Required: true,
		Action: func(ctx *cli.Context, address string) (err error) {
			if !common.IsHexAddress(address) {
				err = fmt.Errorf("wrong value for --pairAddr: %v, must be a hexadecimal address", PairAddress)
				return
			}
			TokenAddress = common.HexToAddress(address)
			return
		},
	}

	DecimalsFlag = cli.UintFlag{
		Name:        "decimals",
		Usage:       "Decimals of token",
		Value:       18,
		EnvVars:     []string{"DECIMALS"},
		Destination: &Decimals,
	}
)
