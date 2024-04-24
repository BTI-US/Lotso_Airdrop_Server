package database

import (
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func GetTransactionCount(address string) (transactionCount *model.TransactionCount, err error) {
	transactionCount, err = mysql.GetTransactionCountByAddress(address)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Get TransactionCount from mysql failed, err:", err)
	}

	return
}

func InsertTransactionCount(transactionCount *model.TransactionCount) (ID uint, err error) {
	if len(transactionCount.Address) == 2*(common.AddressLength+1) {
		transactionCount.Address = transactionCount.Address[2:]
	}
	return mysql.InsertTransactionCount(transactionCount)
}

func GetAddressesShouldAirdrop() (transactionCounts *[]model.TransactionCount, err error) {
	return mysql.GetAddressesShouldAirdrop()
}
