package routes

import (
	"Lotso_Airdrop_Server/controller"
	"Lotso_Airdrop_Server/utils/flags"
	"github.com/gin-gonic/gin"
)

func addInfoRoutesV1(rg *gin.RouterGroup) {
	info := rg.Group("/info")
	{
		info.GET("/apply_airdrop", controller.ApplyAirdrop)
		info.GET("/recipients_count", controller.RecipientsCount)
		addInfoDebugRoutesV1(info)
	}
}

func addInfoDebugRoutesV1(rg *gin.RouterGroup) {
	if flags.Debug {
		rg.GET("/addresses_should_airdrop", controller.AddressesShouldAirdrop)
		rg.POST("/distribute_airdrops", controller.DistributeAirdrops)
		rg.POST("/distribute_airdrops_to", controller.DistributeAirdropsTo)
		rg.POST("/claim_airdrop", controller.ClaimAirdrop)
	}
}

func addInfoRoutesV2(rg *gin.RouterGroup) {
	info := rg.Group("/info")
	{
		info.POST("/set_airdrop", controller.SetAirdrop)
		info.POST("/reward_parent", controller.RewardParent)
	}
}
