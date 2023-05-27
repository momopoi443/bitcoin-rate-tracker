package services

import (
	"github.com/gin-gonic/gin"
	"github.com/momopoi443/bitcoin-rate-tracker/emailsRepo"
	log "github.com/sirupsen/logrus"
)

func HandlePostSubscribe(c *gin.Context) {
	email := c.PostForm("email")

	exist, _ := emailsRepo.Exist(email)
	if exist {
		c.AbortWithStatusJSON(409, gin.H{
			"error": "Email already exists",
		})
		return
	}

	err := emailsRepo.Add(email)
	if err != nil {
		log.Error("Помилка при записі до файлу нового імейлу: " + err.Error())
		return
	}

	c.JSON(200, gin.H{
		"message": "Email subscribed",
	})
}
