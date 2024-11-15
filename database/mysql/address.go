package mysql

import "Lotso_Airdrop_Server/model"

func GetAddressByAddressHex(addressHex string) (address *model.Address, err error) {
	err = db.Where("address = ?", addressHex).First(&address).Error
	return
}

func GetAddressById(id uint) (address *model.Address, err error) {
	err = db.Where("id = ?", id).First(&address).Error
	return
}

func InsertAddress(address *model.Address) (ID uint, err error) {
	err = db.Create(address).Error
	ID = address.ID
	return
}
