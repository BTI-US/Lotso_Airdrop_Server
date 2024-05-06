package controller

import (
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/service"
	"Lotso_Airdrop_Server/utils"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

// RecipientInfo 获取空投合约已经调用领取空投的次数和被领取空投的总量
func RecipientInfo(c *gin.Context) {
	c.JSON(http.StatusOK, service.RecipientInfo())
}

// ApplyAirdrop 接收地址参数，向数据库、链上查询此地址的交易数量并存入mysql，如果满足空投条件，则会按计划发放空投
func ApplyAirdrop(c *gin.Context) {
	address, ok := c.GetQuery("address")
	if !ok {
		response := base.NewErrorResponse(errors.New("missing parameter `address`"), base.WrongParams)
		c.JSON(http.StatusOK, response)
		return
	}

	if !common.IsHexAddress(address) {
		response := base.NewErrorResponse(nil, base.InvalidAddress)
		c.JSON(http.StatusOK, response)
		return
	}

	response, cacheControl := service.ApplyAirdrop(common.HexToAddress(address))
	if cacheControl != "" {
		c.Writer.Header().Set("Cache-Control", "public, immutable, max-age="+cacheControl)
	}
	c.JSON(http.StatusOK, response)
}

// DistributeAirdrops 发放mysql所记录的所有需要发放的空投
func DistributeAirdrops(c *gin.Context) {
	c.JSON(http.StatusOK, service.DistributeAirdrops())
}

// AddressesShouldAirdrop 获得mysql所记录的所有需要发放空投的地址
func AddressesShouldAirdrop(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAddressesShouldAirdrop())
}

// DistributeAirdropsTo 接收地址和空投数量，并立刻对其发送空投，不对mysql记录做任何修改
func DistributeAirdropsTo(c *gin.Context) {
	var param model.DistributeParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.WrongParams))
		return
	}

	if !common.IsHexAddress(param.Address) {
		response := base.NewErrorResponse(nil, base.InvalidAddress)
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, service.DistributeAirdropsTo(common.HexToAddress(param.Address), param.Amount.Uint64()))
}

// ClaimAirdrop 接收私钥为参数，调用合约来领取空投
func ClaimAirdrop(c *gin.Context) {
	var param model.ClaimParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.WrongParams))
		return
	}

	if utils.Has0xPrefix(param.PrivateKey) {
		param.PrivateKey = param.PrivateKey[2:]
	}

	privateKey, err := crypto.HexToECDSA(param.PrivateKey)
	if err != nil {
		response := base.NewErrorResponse(err, base.WrongParams)
		c.JSON(http.StatusOK, response)
		return
	}
	c.JSON(http.StatusOK, service.ClaimAirdrop(privateKey))
}

// SetAirdrop 接收地址和空投数量并存入mysql，会按计划发放空投
func SetAirdrop(c *gin.Context) {
	var param model.DistributeParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.WrongParams))
		return
	}

	if !common.IsHexAddress(param.Address) {
		response := base.NewErrorResponse(nil, base.InvalidAddress)
		c.JSON(http.StatusOK, response)
		return
	}

	if param.Amount.Uint64() > 100000000 {
		response := base.NewErrorResponse(errors.New("amount should <= 100000000"), base.WrongParams)
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, service.SetAirdrop(common.HexToAddress(param.Address), param.Amount.Uint64()))
}

func AppendAirdrop(c *gin.Context) {
	var param model.DistributeParam
	err := c.ShouldBind(&param)
	if err != nil {
		c.JSON(http.StatusOK, base.NewErrorResponse(err, base.WrongParams))
		return
	}

	if !common.IsHexAddress(param.Address) {
		response := base.NewErrorResponse(nil, base.InvalidAddress)
		c.JSON(http.StatusOK, response)
		return
	}

	if param.Amount.Uint64() > 100000000 {
		response := base.NewErrorResponse(errors.New("amount should <= 100000000"), base.WrongParams)
		c.JSON(http.StatusOK, response)
		return
	}

	c.JSON(http.StatusOK, service.AppendAirdrop(common.HexToAddress(param.Address), param.Amount.Uint64()))
}

func CheckEligibility(c *gin.Context) {
	address, ok := c.GetQuery("address")
	if !ok {
		response := base.NewErrorResponse(errors.New("missing parameter `address`"), base.WrongParams)
		c.JSON(http.StatusOK, response)
		return
	}

	if !common.IsHexAddress(address) {
		response := base.NewErrorResponse(nil, base.InvalidAddress)
		c.JSON(http.StatusOK, response)
		return
	}

	isBuyer, err := service.IsBuyer(common.HexToAddress(address))
	if err != nil {
		response := base.NewErrorResponse(err, base.GetTransactionTopicFailed)
		c.JSON(http.StatusOK, response)
		return
	}

	response := base.NewDataResponse(isBuyer)
	c.JSON(http.StatusOK, response)
	return
}
