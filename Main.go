package main

import (
	"github.com/momopoi443/bitcoin-rate-tracker/controllers"
	log "github.com/sirupsen/logrus"
	"os"
)

func SetupLogger() {
	// Встановлення рівня логу
	log.SetLevel(log.InfoLevel)

	// Встановлення формату виводу логів
	log.SetFormatter(&log.TextFormatter{
		FullTimestamp: true,
	})

	// Встановлення виводу в консоль
	log.SetOutput(os.Stdout)
}

func main() {
	SetupLogger()
	router := controllers.SetupRouter()
	router.Run()
}
