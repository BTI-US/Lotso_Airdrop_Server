package database

import (
	"Lotso_Airdrop_Server/database/mysql"
)

func Setup() (err error) {
	err = mysql.Connect()
	return
}
