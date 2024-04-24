package routes

import (
	"Lotso_Airdrop_Server/utils/flags"
	"github.com/gin-gonic/gin"
	"github.com/labstack/gommon/log"
	"github.com/unrolled/secure"
)

var (
	router = gin.Default()
)

func TlsHandler(port string) gin.HandlerFunc {
	return func(c *gin.Context) {
		secureMiddleware := secure.New(secure.Options{
			SSLRedirect: true,
			SSLHost:     ":" + port,
		})
		err := secureMiddleware.Process(c.Writer, c.Request)

		// If there was an error, do not continue.
		if err != nil {
			return
		}

		c.Next()
	}
}

// Run will start the server
func Run(port string) {
	getRoutes()
	if len(flags.SslCertPath) != 0 && len(flags.SslKeyPath) != 0 {
		router.Use(TlsHandler(port))
		log.Info("SSL is enabled.")
		log.Fatal(router.RunTLS(":"+port, flags.SslCertPath, flags.SslKeyPath))
	} else {
		log.Info("SSL is not enabled.")
		log.Fatal(router.Run(":" + port))
	}
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addInfoRoutesV1(v1)

	v2 := router.Group("/v2")
	addInfoRoutesV2(v2)
}
