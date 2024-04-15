package main

import (
	"Lotso_Airdrop_Server/cron"
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/routes"
	"Lotso_Airdrop_Server/utils"
	"Lotso_Airdrop_Server/utils/flags"
	"fmt"
	"github.com/labstack/gommon/log"
	"gopkg.in/urfave/cli.v1"
	"os"
)

const (
	clientIdentifier = "Lotso Airdrop Server"
	clientVersion    = "1.0.0"
	clientUsage      = "Lotso Airdrop Server"
)

var (
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		flags.PortFlag,
		flags.SslCertFlag,
		flags.SslKeyFlag,
		flags.ApiUrlFlag,
		flags.CutoffBlockFlag,
		flags.ContractAddressFlag,
		flags.PrivateKeyFlag,
		flags.ChainIDFlag,
		flags.DebugFlag,
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
		err = fmt.Errorf("invalid cutoff block")
		return
	}
	if !utils.Has0xPrefix(flags.CutoffBlock) {
		flags.CutoffBlock = "0x" + flags.CutoffBlock
	}
	if utils.Has0xPrefix(flags.PrivateKey) {
		flags.PrivateKey = flags.PrivateKey[2:]
	}
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
