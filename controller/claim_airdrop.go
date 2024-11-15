package controller

import (
	"Lotso_Airdrop_Server/model"
	"Lotso_Airdrop_Server/model/base"
	"Lotso_Airdrop_Server/service"
	"Lotso_Airdrop_Server/utils"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/gin-gonic/gin"
	"net/http"
)

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
	if err != nil || !common.IsHexAddress(param.ContractAddress) {
		response := base.NewErrorResponse(err, base.WrongParams)
		c.JSON(http.StatusOK, response)
		return
	}

	address := common.HexToAddress(param.ContractAddress)

	proofs := make([][32]byte, len(param.Proofs))
	for i := 0; i < len(param.Proofs); i++ {
		proofs[i] = [32]byte(common.Hex2Bytes(param.Proofs[i]))
	}

	c.JSON(http.StatusOK, service.ClaimAirdrop(param.AirdropCount, address, proofs, privateKey))
}
