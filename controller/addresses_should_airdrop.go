package controller

import (
	"Lotso_Airdrop_Server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// AddressesShouldAirdrop 获得mysql所记录的所有需要发放空投的地址
func AddressesShouldAirdrop(c *gin.Context) {
	c.JSON(http.StatusOK, service.GetAddressesShouldAirdrop())
}
