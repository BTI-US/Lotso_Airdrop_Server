package service

import (
	"Lotso_Airdrop_Server/database"
	"Lotso_Airdrop_Server/database/mysql"
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/utils"
	"crypto/ecdsa"
	"github.com/ethereum/go-ethereum/common"
	"github.com/labstack/gommon/log"
	"math/big"
	"time"
)

var AirdropCount, _ = new(big.Int).SetString("100000000000000000000000", 10)

func GetTransactionCount(address common.Address) (response *base.Response, cacheControl string) {
	// With 0x prefix
	addressHex := address.Hex()

	transactionCount, err := database.GetTransactionCount(addressHex[2:])
	if err != nil && err.Error() != "record not found" {
		return base.NewErrorResponse(err, base.GetTransactionCountFailed), cacheControl
	}
	transactionCount.AirdropCount = new(big.Int).SetInt64(0)
	if err == nil {
		cacheControl = "600"
		transactionCount.Address = "0x" + transactionCount.Address
		if transactionCount.TransactionCount > 10 {
			transactionCount.AirdropCount = AirdropCount
		}
		response = base.NewDataResponse(transactionCount)
		return
	}

	transactionCountUInt64, err := EthGetTransactionCount(addressHex)
	if err != nil {
		return base.NewErrorResponse(err, base.GetTransactionCountFailed), cacheControl
	}
	transactionCount.Address = addressHex
	transactionCount.TransactionCount = transactionCountUInt64
	transactionCount.ScheduledDelivery = utils.ZeroTime()
	if transactionCount.TransactionCount > 10 {
		transactionCount.AirdropCount = AirdropCount
		transactionCount.ScheduledDelivery = utils.NextOddHourTime()
	}

	_, err = database.InsertTransactionCount(transactionCount)
	if err != nil {
		log.Errorf("InsertTransactionCount Error: %v", err.Error())
	}

	if len(transactionCount.Address) == 2*(common.AddressLength) {
		transactionCount.Address = "0x" + transactionCount.Address
	}

	cacheControl = "600"
	response = base.NewDataResponse(transactionCount)
	return
}

func DistributeAirdrops() (response *base.Response) {
	// Before distributing, we need to remove duplicate addresses
	_ = mysql.RemoveDuplicateAddresses()

	transactionCounts, err := database.GetAddressesShouldAirdrop()
	if err != nil {
		return base.NewErrorResponse(err, base.GetTransactionCountFailed)
	}

	if len(*transactionCounts) == 0 {
		response = base.NewResponse()
		response.Message = "no airdrops need to be distributed"
		return
	}

	var result [][]model.TransactionCount

	for i := 0; i < len(*transactionCounts); i += 1000 {
		end := i + 1000
		if end > len(*transactionCounts) {
			end = len(*transactionCounts)
		}
		result = append(result, (*transactionCounts)[i:end])
	}
	var hashes []string

	for _, subSlice := range result {
		var hash common.Hash
		hash, err = LotsoDistributeAirdrops(&subSlice, AirdropCount)
		if err != nil {
			log.Error("LotsoDistributeAirdrops error: ", err)
			hashes = append(hashes, "")
		} else {
			hashes = append(hashes, hash.Hex())
			_ = mysql.SetTransactionCountsAsHasAirdropped(transactionCounts)
		}
	}

	response = base.NewDataResponse(hashes)
	return
}

func GetAddressesShouldAirdrop() (response *base.Response) {
	transactionCounts, err := database.GetAddressesShouldAirdrop()
	if err != nil {
		return base.NewErrorResponse(err, base.GetTransactionCountFailed)
	}

	response = base.NewDataResponse(transactionCounts)
	return
}

func DistributeAirdropsTo(address common.Address, amount *big.Int) (response *base.Response) {
	var transactionCounts []model.TransactionCount
	airdrop := model.NewAirdrop(&address, amount)
	transactionCounts = append(transactionCounts, *airdrop)
	hash, err := LotsoDistributeAirdrops(&transactionCounts, amount)
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
