package flags

import "gopkg.in/urfave/cli.v1"

var (
	PortFlag = cli.StringFlag{
		Name:   "port, p",
		Usage:  "Server port",
		Value:  "1423",
		EnvVar: "SERVER_PORT",
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
		Usage:       "API URL",
		EnvVar:      "API_URL",
		Destination: &ApiUrl,
	}
	PrivateKeyFlag = cli.StringFlag{
		Name:        "private, pri",
		Usage:       "Private key hex string",
		EnvVar:      "PRIVATE_KEY",
		Destination: &PrivateKey,
	}
	ChainIDFlag = cli.Int64Flag{
		Name:        "chain_id",
		Usage:       "ChainID",
		EnvVar:      "CHAIN_ID",
		Destination: &ChainID,
	}

	CutoffBlockFlag = cli.StringFlag{
		Name:        "CutoffBlock",
		Usage:       "Cutoff block",
		EnvVar:      "CUTOFF_BLOCK",
		Value:       "latest",
		Destination: &CutoffBlock,
	}

	ContractAddressFlag = cli.StringFlag{
		Name:        "contract",
		Usage:       "Contract address",
		EnvVar:      "CONTRACT_ADDRESS",
		Destination: &ContractAddress,
	}

	MysqlHostFlag = cli.StringFlag{
		Name:        "MysqlHost",
		Usage:       "MysqlHost",
		Value:       "127.0.0.1",
		EnvVar:      "MYSQL_HOST",
		Destination: &MysqlHost,
	}
	MysqlPortFlag = cli.IntFlag{
		Name:        "MysqlPort",
		Usage:       "MysqlPort",
		Value:       3306,
		EnvVar:      "MYSQL_PORT",
		Destination: &MysqlPort,
	}
	MysqlUserFlag = cli.StringFlag{
		Name:        "MysqlUser",
		Usage:       "MysqlUser",
		Value:       "root",
		EnvVar:      "MYSQL_USER",
		Destination: &MysqlUser,
	}
	MysqlPasswdFlag = cli.StringFlag{
		Name:        "MysqlPasswd",
		Usage:       "MysqlPasswd",
		Value:       "123456",
		EnvVar:      "MYSQL_PASSWD",
		Destination: &MysqlPasswd,
	}
	MysqlDBFlag = cli.StringFlag{
		Name:        "MysqlDB",
		Usage:       "mysqlDatabase",
		Value:       "lotso",
		EnvVar:      "MYSQL_DB",
		Destination: &MysqlDB,
	}
)
