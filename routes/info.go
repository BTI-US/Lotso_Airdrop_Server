package routes

import (
	"Lotso_Airdrop_Server/controller"
	"Lotso_Airdrop_Server/utils/flags"
	"github.com/gin-gonic/gin"
)

func addInfoRoutesV1(rg *gin.RouterGroup) {
	info := rg.Group("/info")
	{
		info.GET("/transaction_count", controller.TransactionCount)
		if flags.Debug {
			info.GET("/addresses_should_airdrop", controller.AddressesShouldAirdrop)
			info.POST("/distribute_airdrops", controller.DistributeAirdrops)
		}
	}
}
