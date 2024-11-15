package routes

import (
	mingfuflags "github.com/Fu-XDU/mingfu_go_common/flags"
	mingfuroutes "github.com/Fu-XDU/mingfu_go_common/routes"
	"github.com/gin-gonic/gin"
)

// Run will start the server
func Run() {
	mingfuroutes.Run(mingfuflags.ServerPort, mingfuflags.SslCertPath, mingfuflags.SslKeyPath, mingfuflags.TrustedProxies, getRoutes)
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes(router *gin.Engine) {
	v1 := router.Group("/v1")
	addAirdropRoutesV1(v1)
}
