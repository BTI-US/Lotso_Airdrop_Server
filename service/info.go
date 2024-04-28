package service

import (
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/utils"
	"Lotso_Airdrop_Server/utils/flags"
	"crypto/ecdsa"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"gorm.io/gorm"
	"math/big"
	"time"
)

var AirdropCount uint64 = 100000

func ApplyAirdrop(address common.Address) (response *base.Response, cacheControl string) {
	// With 0x prefix
	addressHex := address.Hex()

	airdropItem, err := database.GetAirdropItem(addressHex[2:])
	if err != nil && !errors.Is(err, gorm.ErrRecordNotFound) {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed), cacheControl
	}

	if err == nil {
		cacheControl = "600"
		airdropItem.Address = "0x" + airdropItem.Address
		if airdropItem.TransactionCount > 10 {
			airdropItem.AirdropCount = AirdropCount
		}
		response = base.NewDataResponse(airdropItem)
		return
	}

	transactionCountUInt64, err := EthGetTransactionCount(addressHex)
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed), cacheControl
	}
	airdropItem.Address = addressHex
	airdropItem.TransactionCount = transactionCountUInt64
	airdropItem.ScheduledDelivery = utils.ZeroTime()
	if airdropItem.TransactionCount > 10 {
		airdropItem.AirdropCount = AirdropCount
		airdropItem.ScheduledDelivery = utils.NextOddHourTime()
	}

	_, err = database.InsertAirdropItem(airdropItem)
	if err != nil {
		log.Errorf("InsertAirdropItem Error: %v", err.Error())
	}

	if len(airdropItem.Address) == 2*(common.AddressLength) {
		airdropItem.Address = "0x" + airdropItem.Address
	}

	cacheControl = "600"
	response = base.NewDataResponse(airdropItem)
	return
}

func DistributeAirdrops() (response *base.Response) {
	airdropItems, err := database.GetAddressesShouldAirdrop()
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed)
	}

	if len(*airdropItems) == 0 {
		response = base.NewResponse()
		response.Message = "no airdrops need to be distributed"
		return
	}

	var result [][]model.AirdropItem

	for i := 0; i < len(*airdropItems); i += 1000 {
		end := i + 1000
		if end > len(*airdropItems) {
			end = len(*airdropItems)
		}
		result = append(result, (*airdropItems)[i:end])
	}
	var hashes []string

	for _, subSlice := range result {
		var hash common.Hash

		hash, err = LotsoDistributeAirdrops(&subSlice)
		if err != nil {
			log.Error("LotsoDistributeAirdrops error: ", err)
			hashes = append(hashes, "")
		} else {
			hashes = append(hashes, hash.Hex())
			_ = mysql.SetAirdropItemsAsHasAirdropped(airdropItems)
		}
	}

	response = base.NewDataResponse(hashes)
	return
}

func GetAddressesShouldAirdrop() (response *base.Response) {
	airdropItems, err := database.GetAddressesShouldAirdrop()
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed)
	}

	response = base.NewDataResponse(airdropItems)
	return
}

func DistributeAirdropsTo(address common.Address, amount uint64) (response *base.Response) {
	var airdropItems []model.AirdropItem
	airdrop := model.NewAirdrop(&address, amount)
	airdropItems = append(airdropItems, *airdrop)
	hash, err := LotsoDistributeAirdrops(&airdropItems)
	if err != nil {
		response = base.NewErrorResponse(err, base.DistributeAirdropsFailed)
		return
	}
	response = base.NewDataResponse(hash)
	return
}

func ClaimAirdrop(privateKey *ecdsa.PrivateKey) (response *base.Response) {
	hash, err := LotsoClaimAirdrops(privateKey)
	if err != nil {
		response = base.NewErrorResponse(err, base.ClaimAirdropsFailed)
		return
	}

	response = base.NewDataResponse(hash)
	return
}

var (
	recipientsCount, _        = new(big.Int).SetString("0", 10)
	recipientsCountLastUpdate int64
)

func RecipientsCount() (response *base.Response) {
	now := time.Now().Unix()
	if now-recipientsCountLastUpdate < 30 {
		response = base.NewDataResponse(recipientsCount)
		return
	}

	recipientsCountLastUpdate = now
	count, err := LotsoRecipientsCount()
	if err != nil {
		response = base.NewErrorResponse(err, base.RecipientsCountFailed)
		response.Data = recipientsCount
		return
	}
	recipientsCount = count

	response = base.NewDataResponse(count)
	return
}

func SetAirdrop(address common.Address, amount uint64) (response *base.Response) {
	// With 0x prefix
	addressHex := address.Hex()

	airdropItem, err := database.GetAirdropItem(addressHex[2:])
	if err == nil {
		airdropItem.Address = "0x" + airdropItem.Address
		response = base.NewErrorResponse(err, base.AirdropItemAlreadyExists)
		response.Data = airdropItem
		return response
	}
	if !errors.Is(err, gorm.ErrRecordNotFound) {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed)
	}
	airdropItem.AirdropCount = amount
	airdropItem.Address = addressHex[2:]
	airdropItem.ScheduledDelivery = utils.NextOddHourTime()
	_, err = database.InsertAirdropItem(airdropItem)
	if err != nil {
		return base.NewErrorResponse(err, base.SaveAirdropItemFailed)
	}
	response = base.NewDataResponse(airdropItem)
	return
}

func AppendAirdrop(address common.Address, amount uint64) (response *base.Response) {
	addressHex := address.Hex()

	item, err := mysql.GetAirdropItemByAddress(address.Hex())
	if err != nil {
		return base.NewErrorResponse(err, base.GetAirdropItemFailed)
	}

	isBuyer, err := IsBuyer(address)
	if err != nil {
		return base.NewErrorResponse(err, base.GetTransactionTopicFailed)
	}

	newAirdropCount := item.AirdropCount + amount

	if (isBuyer && newAirdropCount > flags.BuyerRewardLimit) || (!isBuyer && newAirdropCount > flags.NotBuyerRewardLimit) {
		response = base.NewErrorResponse(nil, base.RewardLimitReached)
		response.Data = item
		return
	}

	item, err = mysql.AppendAirdropAmount(addressHex[2:], amount)
	if err != nil {
		return base.NewErrorResponse(err, base.SaveAirdropItemFailed)
	}
	return base.NewDataResponse(item)
}
