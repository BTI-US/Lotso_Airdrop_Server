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

func RecipientsCount(c *gin.Context) {
	c.JSON(http.StatusOK, service.RecipientsCount())
}

func TransactionCount(c *gin.Context) {
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

	response, cacheControl := service.GetTransactionCount(common.HexToAddress(address))
	if cacheControl != "" {
		c.Writer.Header().Set("Cache-Control", "public, immutable, max-age="+cacheControl)
	}
	c.JSON(http.StatusOK, response)
}

func DistributeAirdrops(c *gin.Context) {
	c.JSON(http.StatusOK, service.DistributeAirdrops())
}

func AddressesShouldAirdrop(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAddressesShouldAirdrop())
}

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
	c.JSON(http.StatusOK, service.DistributeAirdropsTo(common.HexToAddress(param.Address), &param.Amount))
}

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
