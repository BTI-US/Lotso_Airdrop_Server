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
		info.GET("/recipients_count", controller.RecipientsCount)
		if flags.Debug {
			info.GET("/addresses_should_airdrop", controller.AddressesShouldAirdrop)
			info.POST("/distribute_airdrops", controller.DistributeAirdrops)
			info.POST("/distribute_airdrops_to", controller.DistributeAirdropsTo)
			info.POST("/claim_airdrop", controller.ClaimAirdrop)
		}
	}
}
