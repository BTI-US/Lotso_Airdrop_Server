package database

import (
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

func GetAirdrops(address string) (airdrops *[]model.Airdrop, err error) {
	airdrops, err = mysql.GetAirdropsByAddress(address)
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		log.Error("Get Airdrop from mysql failed, err:", err)
	}

	return
}

func InsertAirdrop(airdrop *model.Airdrop) (ID uint, err error) {
	if len(airdrop.Address.Address) == 2*(common.AddressLength+1) {
		airdrop.Address.Address = airdrop.Address.Address[2:]
	}
	return mysql.InsertAirdrop(airdrop)
}

func GetAddressesShouldAirdrop() (airdrops *[]model.Airdrop, err error) {
	return mysql.GetAirdropsShouldAirdrop()
}

func UpdateAirdrops(airdropItems *[]model.Airdrop) (err error) {
	return mysql.UpdateAirdrops(airdropItems)
}
