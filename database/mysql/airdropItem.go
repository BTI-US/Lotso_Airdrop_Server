package mysql

import (
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/utils"
)

func InsertAirdropItem(airdropItem *model.AirdropItem) (ID uint, err error) {
	err = db.Create(airdropItem).Error
	ID = airdropItem.ID
	return
}

func GetAirdropItemByID(ID int) (airdropItem *model.AirdropItem, err error) {
	err = db.Where("id = ?", ID).First(&airdropItem).Error
	return
}

func GetAirdropItemByAddress(address string) (airdropItem *model.AirdropItem, err error) {
	err = db.Where("address = ?", address).First(&airdropItem).Error
	return
}

func DelAirdropItemByID(ID int) (err error) {
	err = db.Delete(model.AirdropItem{}, ID).Error
	return
}

func GetAddressesShouldAirdrop() (airdropItems *[]model.AirdropItem, err error) {
	selectCol := []string{"address", "airdrop_count", "has_airdropped_count", "scheduled_delivery"}
	err = db.Where("airdrop_count > has_airdropped_count AND scheduled_delivery > ?", utils.ZeroTime()).Select(selectCol).Find(&airdropItems).Error
	return
}

func SetAirdropItemsAsHasAirdropped(airdropItems *[]model.AirdropItem) (err error) {
	for _, tc := range *airdropItems {
		db.Model(&model.AirdropItem{}).Where("address = ?", tc.Address).Update("has_airdropped_count", tc.AirdropCount)
	}
	return
}

func RemoveDuplicateAddresses() (err error) {
	err = db.Exec("DELETE FROM airdrop_items WHERE id NOT IN (SELECT tc.min_id FROM (SELECT MIN(id) AS min_id FROM transaction_counts GROUP BY address) tc)").Error
	return
}
