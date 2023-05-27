package services

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/momopoi443/bitcoin-rate-tracker/emailsRepo"
	"gopkg.in/gomail.v2"
)

func HandlePostSendEmails(c *gin.Context) {
	rate := FetchBitcoinRate()
	subscribedEmails, _ := emailsRepo.ListAll()

	for _, email := range subscribedEmails {
		sendEmail(
			"example@gmail.com",
			email,
			"курс біткоіну (BTC) у гривні (UAH)",
			"BTC до UAH: "+fmt.Sprintf("%f", rate))
	}

	c.JSON(200, gin.H{
		"message": "Emails sent",
	})
}

func sendEmail(from, to, subject, body string) error {
	// Налаштування локального SMTP сервера
	smtpHost := "localhost"
	smtpPort := 25

	// Формування повідомлення
	mailer := gomail.NewMessage()
	mailer.SetHeader("From", from)
	mailer.SetHeader("To", to)
	mailer.SetHeader("Subject", subject)
	mailer.SetBody("text/plain", body)

	// Відправлення електронного листа
	dialer := gomail.Dialer{Host: smtpHost, Port: smtpPort}
	return dialer.DialAndSend(mailer)
}
