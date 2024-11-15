package service

import (
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/utils"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
)

type ApplyAirdropResponse struct {
	Address  string           `json:"address"`
	Airdrops *[]model.Airdrop `json:"airdrops"`
}

var AirdropCount uint64 = 100000

// ApplyAirdrop 接收地址参数，向数据库、链上查询此地址的交易数量并存入mysql
// 返回此地址的所有空投信息
func ApplyAirdrop(address common.Address) (response *base.Response, cacheControl string) {
	// With 0x prefix
	addressHex := address.Hex()

	addressItem, err := database.GetAddress(addressHex[2:])
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		response = base.NewErrorResponse(err, base.GetAddressFailed)
		return
	}

	// 数据库中找到此地址，返回此地址所有空投
	if err == nil {
		addressItem.Address = addressHex

		airdrops, err := getAirdropsByAddress(address)
		if err != nil {
			response = base.NewErrorResponse(err, base.UnknownError)
			return
		}

		airdropsResponse := &ApplyAirdropResponse{
			Address:  addressHex,
			Airdrops: airdrops,
		}

		response = base.NewDataResponse(airdropsResponse)
		cacheControl = "600"
		return
	}

	// 数据库中未找到此地址，新增
	transactionCountUInt64, err := EthGetTransactionCount(addressHex)
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed), cacheControl
	}

	id, err := database.InsertAddress(addressHex[2:], transactionCountUInt64)
	if err != nil {
		return base.NewErrorResponse(err, base.SaveAirdropItemFailed), cacheControl
	}

	if transactionCountUInt64 > 10 {
		airdropItem := &model.Airdrop{
			AddressID:         id,
			AirdropCount:      AirdropCount,
			ContractAddress:   "",
			ScheduledDelivery: utils.NextOddHourTime(),
		}

		_, err = database.InsertAirdrop(airdropItem)
		if err != nil {
			log.Errorf("InsertAirdropItem Error: %v", err.Error())
			response = base.NewErrorResponse(err, base.SaveAirdropItemFailed)
			return
		}
	}

	airdrops, err := getAirdropsByAddress(address)
	if err != nil {
		response = base.NewErrorResponse(err, base.UnknownError)
		return
	}

	airdropsResponse := &ApplyAirdropResponse{
		Address:  addressHex,
		Airdrops: airdrops,
	}

	response = base.NewDataResponse(airdropsResponse)
	cacheControl = "600"
	return
}

func getAirdropsByAddress(address common.Address) (airdrops *[]model.Airdrop, err error) {
	addressHex := address.Hex()[2:]
	airdrops, err = database.GetAirdrops(addressHex)
	if err != nil {
		return
	}
	return
}
