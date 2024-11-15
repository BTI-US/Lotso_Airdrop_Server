package controller

import (
	"Lotso_Airdrop_Server/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

// DistributeAirdrops 发放mysql所记录的所有需要发放的空投
func DistributeAirdrops(c *gin.Context) {
	c.JSON(http.StatusOK, service.DistributeAirdrops())
}
