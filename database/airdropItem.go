package database

import (
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func GetAirdropItem(address string) (airdropItem *model.AirdropItem, err error) {
	airdropItem, err = mysql.GetAirdropItemByAddress(address)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Get AirdropItem from mysql failed, err:", err)
	}

	return
}

func InsertAirdropItem(airdropItem *model.AirdropItem) (ID uint, err error) {
	if len(airdropItem.Address) == 2*(common.AddressLength+1) {
		airdropItem.Address = airdropItem.Address[2:]
	}
	return mysql.InsertAirdropItem(airdropItem)
}

func GetAddressesShouldAirdrop() (airdropItems *[]model.AirdropItem, err error) {
	return mysql.GetAddressesShouldAirdrop()
}
