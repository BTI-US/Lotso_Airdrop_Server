package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/unrolled/secure"
	"log"
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

		if err != nil {
			return
		}

		c.Next()
	}
}

// Run will start the server
func Run(port, sslCertPath, sslKeyPath string, trustedProxies []string) {
	if len(trustedProxies) != 0 {
		router.ForwardedByClientIP = true
		err := router.SetTrustedProxies(trustedProxies)
		if err != nil {
			log.Fatalf("Wrong TrustedProxies, %v", err)
		}
	}
	getRoutes()
	if len(sslCertPath) != 0 && len(sslKeyPath) != 0 {
		log.Println("SSL is enabled.")
		router.Use(TlsHandler(port))
		log.Fatal(router.RunTLS(":"+port, sslCertPath, sslKeyPath))
	} else {
		log.Println("SSL is not enabled.")
		log.Fatal(router.Run(":" + port))
	}
}

// getRoutes will create our routes of our entire application
// this way every group of routes can be defined in their own file
// so this one won't be so messy
func getRoutes() {
	v1 := router.Group("/v1")
	addInfoRoutesV1(v1)
}
