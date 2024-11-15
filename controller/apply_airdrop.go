package controller

import (
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/service"
	"errors"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
