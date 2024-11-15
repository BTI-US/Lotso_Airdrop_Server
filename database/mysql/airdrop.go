package mysql

import (
	"Lotso_Airdrop_Server/model"
)

func InsertAirdrop(airdropItem *model.Airdrop) (ID uint, err error) {
	err = db.Create(airdropItem).Error
	ID = airdropItem.ID
	return
}

func GetAirdropByID(ID int) (airdropItem *model.Airdrop, err error) {
	err = db.Where("id = ?", ID).First(&airdropItem).Error
	return
}

func GetAirdropsByAddress(address string) (airdropItems *[]model.Airdrop, err error) {
	addr, err := GetAddressByAddressHex(address)
	if err != nil {
		return
	}

	err = db.Where("address_id = ?", addr.ID).Find(&airdropItems).Error
	return
}

func DelAirdropByID(ID int) (err error) {
	err = db.Delete(model.Airdrop{}, ID).Error
	return
}

func AirdropExist(address string) (exist bool, err error) {
	var count int64
	err = db.Model(&model.Airdrop{}).Where("address = ?", address).Count(&count).Error
	exist = count > 0
	return
}

func GetAirdropsShouldAirdrop() (airdropItems *[]model.Airdrop, err error) {
	err = db.Preload("Address").Where("contract_address = ?", "").Find(&airdropItems).Error
	return
}

func UpdateAirdrops(airdropItems *[]model.Airdrop) (err error) {
	for _, tc := range *airdropItems {
		db.Save(&tc)
	}
	return
}
