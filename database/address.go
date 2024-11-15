package database

import (
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"errors"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func GetAddress(addressHex string) (address *model.Address, err error) {
	address, err = mysql.GetAddressByAddressHex(addressHex)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Get AirdropItem from mysql failed, err:", err)
	}
	return
}

func InsertAddress(addressHex string, transactionCount uint64) (id uint, err error) {
	address := &model.Address{
		Address:          addressHex,
		TransactionCount: transactionCount,
	}

	id, err = mysql.InsertAddress(address)
	if err != nil {
		return
	}
	return
}
