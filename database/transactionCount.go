package database

import (
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
)

func GetTransactionCount(address string) (transactionCount *model.TransactionCount, err error) {
	transactionCount, err = mysql.GetTransactionCountByAddress(address)
	if err != nil && err.Error() != "record not found" {
		log.Error("Get TransactionCount from mysql failed, err:", err)
	}

	return
}

func InsertTransactionCount(transactionCount *model.TransactionCount) (ID int, err error) {
	if len(transactionCount.Address) == 2*(common.AddressLength+1) {
		transactionCount.Address = transactionCount.Address[2:]
	}
	return mysql.InsertTransactionCount(transactionCount)
}

func GetAddressesShouldAirdrop() (transactionCounts *[]model.TransactionCount, err error) {
	return mysql.GetAddressesShouldAirdrop()
}
