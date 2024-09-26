package router

import (
	"atendaslock4go/paths"
	"github.com/gin-gonic/gin"
)

func Initialize() {

	//gin.SetMode(gin.ReleaseMode)
	router := gin.Default()
	err := router.SetTrustedProxies(nil)
	if err != nil {
		return
	}

	// Initialize Routes
	initializeRoutes(router)
	router.RunTLS(":8181", paths.SslCertFile(), paths.SslKeyFile())
}
