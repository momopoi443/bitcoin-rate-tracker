package services

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"net/http"
)

type CoinGeckoResponse struct {
	Bitcoin map[string]float64 `json:"bitcoin"`
}

func HandleGetBitcoinRate(ginContext *gin.Context) {
	bitcoinRate := FetchBitcoinRate()

	ginContext.JSON(200, gin.H{
		"rate": bitcoinRate,
	})
}

func FetchBitcoinRate() float64 {
	var rate float64

	response, err := http.Get("https://api.coingecko.com/api/v3/simple/price?ids=bitcoin&vs_currencies=uah")
	if err != nil {
		log.Error("Помилка при виконанні запиту: " + err.Error())
		return -1
	}

	defer response.Body.Close()

	data, err := ioutil.ReadAll(response.Body)
	if err != nil {
		log.Error("Помилка при читанні відповіді: " + err.Error())
		return -1
	}

	// Розпаковуємо JSON-дані у структуру CoinGeckoResponse
	var result CoinGeckoResponse
	err = json.Unmarshal(data, &result)
	if err != nil {
		log.Error("Помилка при розпаковуванні JSON: " + err.Error())
		return -1
	}

	// Отримуємо курс BTC до UAH
	rate = result.Bitcoin["uah"]

	return rate
}
