package mysql

import (
	"Lotso_Airdrop_Server/model"
	mingfumysql "github.com/Fu-XDU/mingfu_go_common/database/mysql"
	"gorm.io/gorm"
)

var db *gorm.DB

func Connect() (err error) {
	db, err = mingfumysql.Connect(mingfumysql.NewConnOptionsFromFlags(), nil, initMysql)
	return
}

func initMysql(db *gorm.DB) (err error) {
	err = db.AutoMigrate(&model.Address{}, &model.Airdrop{}, &model.TransferTopic{})
	return
}
