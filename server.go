package main

import (
	"Lotso_Airdrop_Server/cron"
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/routes"
	"Lotso_Airdrop_Server/utils"
	"Lotso_Airdrop_Server/utils/flags"
	"fmt"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gopkg.in/urfave/cli.v1"
	"os"
)

const (
	clientIdentifier = "Lotso Airdrop Server"
	clientVersion    = "1.1.1"
	clientUsage      = "Lotso Airdrop Server"
)

var (
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		flags.PortFlag,
		flags.DebugFlag,
		flags.ApiModeFlag,
		flags.BuyerRewardLimitFlag,
		flags.NotBuyerRewardLimitFlag,
	}
	chainFlags = []cli.Flag{
		flags.ApiUrlFlag,
		flags.CutoffBlockFlag,
		flags.ContractAddressFlag,
		flags.PrivateKeyFlag,
		flags.ChainIDFlag,
		flags.DecimalsFlag,
		flags.PairAddressFlag,
	}
	sslFlags = []cli.Flag{
		flags.SslCertFlag,
		flags.SslKeyFlag,
	}
	mysqlFlags = []cli.Flag{
		flags.MysqlHostFlag,
		flags.MysqlPortFlag,
		flags.MysqlUserFlag,
		flags.MysqlPasswdFlag,
		flags.MysqlDBFlag,
	}
)

func init() {
	app.Action = ServerApp
	app.Name = clientIdentifier
	app.Version = clientVersion
	app.Usage = clientUsage
	app.Commands = []cli.Command{}
	app.Flags = append(app.Flags, baseFlags...)
	app.Flags = append(app.Flags, chainFlags...)
	app.Flags = append(app.Flags, sslFlags...)
	app.Flags = append(app.Flags, mysqlFlags...)
}

func ServerApp(ctx *cli.Context) error {
	if args := ctx.Args(); len(args) > 0 {
		return fmt.Errorf("invalid command: %q", args[0])
	}
	err := prepare(ctx)
	if err != nil {
		log.Error(err)
	}
	return err
}

func prepare(ctx *cli.Context) (err error) {
	if !flags.IsValidCutoffBlock(flags.CutoffBlock) {
		err = fmt.Errorf("invalid cutoff block: %v", flags.CutoffBlock)
		return
	}
	if !common.IsHexAddress(flags.ContractAddress) {
		err = fmt.Errorf("invalid contract address: %v", flags.ContractAddress)
		return
	}
	flags.Contract = common.HexToAddress(flags.ContractAddress)
	if !common.IsHexAddress(flags.PairAddress) {
		err = fmt.Errorf("invalid pair address: %v", flags.PairAddress)
		return
	}
	flags.PairAddress = utils.Remove0xPrefix(flags.PairAddress)
	if !utils.Has0xPrefix(flags.CutoffBlock) {
		flags.CutoffBlock = "0x" + flags.CutoffBlock
	}
	flags.PrivateKey = utils.Remove0xPrefix(flags.PrivateKey)
	if err = database.Setup(); err != nil {
		return
	}
	cron.StartWorker("0 1-23/2 * * *") // 0 1-23/2 * * *
	p := ctx.String("port")
	routes.Run(p)
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
