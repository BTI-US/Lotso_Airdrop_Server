package main

import (
	"Lotso_Airdrop_Server/cron"
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/routes"
	"Lotso_Airdrop_Server/utils/flags"
	"fmt"
	mingfuflags "github.com/Fu-XDU/mingfu_go_common/flags"
	"github.com/labstack/gommon/log"
	"github.com/urfave/cli/v2"
	"os"
)

const (
	clientIdentifier = "Lotso Airdrop Server"
	clientVersion    = "2.0.0"
	clientUsage      = "Lotso Airdrop Server"
)

var (
	app       = cli.NewApp()
	baseFlags = []cli.Flag{
		&flags.DebugFlag,
	}
	chainFlags = []cli.Flag{
		&flags.ApiUrlFlag,
		&flags.CutoffBlockFlag,
		&flags.PrivateKeyFlag,
		&flags.ChainIDFlag,
		&flags.DecimalsFlag,
		&flags.PairAddressFlag,
		&flags.TokenAddressFlag,
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
	app.Flags = append(app.Flags, mingfuflags.GinFlags...)
	app.Flags = append(app.Flags, mingfuflags.MysqlFlags...)
}

func ServerApp(ctx *cli.Context) error {
	if args := ctx.Args(); args.Len() > 0 {
		return fmt.Errorf("invalid command: %q", args.First())
	}
	err := prepare()
	if err != nil {
		log.Error(err)
	}
	return err
}

func prepare() (err error) {
	if err = database.Setup(); err != nil {
		return
	}
	cron.StartWorker("0 1-23/2 * * *") // 0 1-23/2 * * *
	routes.Run()
	return
}

func main() {
	if err := app.Run(os.Args); err != nil {
		_, _ = fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
