package mysql

import (
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/utils"
)

func InsertTransactionCount(transactionCount *model.TransactionCount) (ID uint, err error) {
	err = db.Create(transactionCount).Error
	ID = transactionCount.ID
	return
}

func GetTransactionCountByID(ID int) (transactionCount *model.TransactionCount, err error) {
	err = db.Where("id = ?", ID).First(&transactionCount).Error
	return
}

func GetTransactionCountByAddress(address string) (transactionCount *model.TransactionCount, err error) {
	err = db.Where("address = ?", address).First(&transactionCount).Error
	return
}

func DelTransactionCountByID(ID int) (err error) {
	err = db.Delete(model.TransactionCount{}, ID).Error
	return
}

func GetAddressesShouldAirdrop() (transactionCounts *[]model.TransactionCount, err error) {
	selectCol := []string{"address", "airdrop_count", "scheduled_delivery"}
	err = db.Where("has_airdropped = ? AND scheduled_delivery > ?", false, utils.ZeroTime()).Select(selectCol).Find(&transactionCounts).Error
	return
}

func SetTransactionCountsAsHasAirdropped(transactionCounts *[]model.TransactionCount) (err error) {
	for _, tc := range *transactionCounts {
		db.Model(&model.TransactionCount{}).Where("address = ?", tc.Address).Update("has_airdropped", true)
	}
	return
}

func RemoveDuplicateAddresses() (err error) {
	err = db.Exec("DELETE FROM transaction_counts WHERE id NOT IN (SELECT tc.min_id FROM (SELECT MIN(id) AS min_id FROM transaction_counts GROUP BY address) tc)").Error
	return
}
