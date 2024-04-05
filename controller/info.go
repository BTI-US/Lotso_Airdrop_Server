package controller

import (
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/service"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
