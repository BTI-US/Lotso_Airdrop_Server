package main

import (
	"Lotso_Airdrop_Server/cron"
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/routes"
	"Lotso_Airdrop_Server/utils/flags"
	"fmt"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
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
		&flags.PortFlag,
		&flags.DebugFlag,
		&flags.ApiModeFlag,
		&flags.BuyerRewardLimitFlag,
		&flags.NotBuyerRewardLimitFlag,
		&flags.TrustedProxiesFlag,
	}
	chainFlags = []cli.Flag{
		&flags.ApiUrlFlag,
		&flags.CutoffBlockFlag,
		&flags.ContractAddressFlag,
		&flags.PrivateKeyFlag,
		&flags.ChainIDFlag,
		&flags.DecimalsFlag,
		&flags.PairAddressFlag,
	}
	sslFlags = []cli.Flag{
		&flags.SslCertFlag,
		&flags.SslKeyFlag,
	}
	mysqlFlags = []cli.Flag{
		&flags.MysqlHostFlag,
		&flags.MysqlPortFlag,
		&flags.MysqlUserFlag,
		&flags.MysqlPasswdFlag,
		&flags.MysqlDBFlag,
	}
)

func init() {
	app.Action = ServerApp
	app.Name = clientIdentifier
	app.Version = clientVersion
	app.Usage = clientUsage
	app.Commands = []*cli.Command{}
	app.Flags = append(app.Flags, baseFlags...)
	app.Flags = append(app.Flags, chainFlags...)
	app.Flags = append(app.Flags, sslFlags...)
	app.Flags = append(app.Flags, mysqlFlags...)
}

func ServerApp(ctx *cli.Context) error {
	if args := ctx.Args(); args.Len() > 0 {
		return fmt.Errorf("invalid command: %q", args.First())
	}
	err := prepare(ctx)
	if err != nil {
		log.Error(err)
	}
	return err
}

func prepare(ctx *cli.Context) (err error) {
	flags.ProcessFlags(ctx)
	if err = database.Setup(); err != nil {
		return
	}
	cron.StartWorker("0 1-23/2 * * *") // 0 1-23/2 * * *
	port := ctx.String("port")
	sslCertPath := ctx.String("ssl_cert")
	sslKeyPath := ctx.String("ssl_key")
	trustedProxies := ctx.StringSlice("trusted_proxies")
	routes.Run(port, sslCertPath, sslKeyPath, trustedProxies)
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
