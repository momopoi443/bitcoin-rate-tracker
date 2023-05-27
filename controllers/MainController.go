package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/momopoi443/bitcoin-rate-tracker/services"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/api/rate", services.HandleGetBitcoinRate)
	router.POST("/api/subscribe", services.HandlePostSubscribe)
	router.POST("/api/sendEmails", services.HandlePostSendEmails)

	return router
}
