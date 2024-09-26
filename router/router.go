package router

import (
	"atendaslock4go/handler"
	"atendaslock4go/middleware"
	"github.com/gin-gonic/gin"
)

func initializeRoutes(router *gin.Engine) {
	// TODO: make router for signatures.
	signature := router.Group("/signature")
	{
		signature.POST("/validate", handler.SignatureVerification)

	}

	// Client Routes
	router.POST("/ClientConnect", middleware.ClientConnect, handler.ClientConnect)

	router.POST("/PostTest", handler.PostTest)

	router.POST("/ClientScreen", handler.ClientScreen)

}
