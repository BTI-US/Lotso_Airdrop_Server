package mysql

import "Lotso_Airdrop_Server/model"

func TopicExist(from, to string) (exist bool, err error) {
	var count int64
	err = db.Model(&model.TransferTopic{}).Where("`from` = ?", from).Where("`to` = ?", to).Count(&count).Error
	exist = count > 0
	return
}
