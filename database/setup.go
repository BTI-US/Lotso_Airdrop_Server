package database

import (
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/utils/flags"
)

func Setup() (err error) {
	_, err = mysql.Connect(flags.MysqlHost, flags.MysqlPort, flags.MysqlUser, flags.MysqlPasswd, flags.MysqlDB)
	return
}
