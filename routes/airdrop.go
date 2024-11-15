package routes

import (
	"Lotso_Airdrop_Server/controller"
	"Lotso_Airdrop_Server/utils/flags"
	"github.com/gin-gonic/gin"
)

func addAirdropRoutesV1(rg *gin.RouterGroup) {
	info := rg.Group("/airdrop")
	addAirdropDebugRoutesV1(info)
	info.GET("/apply", controller.ApplyAirdrop)
}

func addAirdropDebugRoutesV1(rg *gin.RouterGroup) {
	if flags.Debug {
		rg.GET("/pending_addresses", controller.AddressesShouldAirdrop)
		rg.POST("/distribute", controller.DistributeAirdrops)
		rg.POST("/multi_distribute_to", controller.DistributeAirdropsTo)
		rg.POST("/claim", controller.ClaimAirdrop)
	}
}
