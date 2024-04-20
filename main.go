package main

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

type CryptoData struct {
	Price float64 `json:"rate"` // Change to float64 to match the numeric type of the API response
}

func getCryptoPrice(crypto string) (float64, error) { // Change return type to float64
	apiKey := os.Getenv("COINAPI_KEY")
	url := fmt.Sprintf("https://rest.coinapi.io/v1/exchangerate/%s/USD", crypto)
	client := &http.Client{}
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("X-CoinAPI-Key", apiKey)
	resp, err := client.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data CryptoData
	err = json.Unmarshal(body, &data)
	if err != nil {
		return 0, err
	}

	return data.Price, nil
}

func main() {
	godotenv.Load()
	router := gin.Default()

	if err := router.SetTrustedProxies([]string{"192.168.0.1"}); err != nil {
		log.Fatalf("Could not set trusted proxies: %v", err)
	}

	router.GET("/price/:crypto", func(c *gin.Context) {
		crypto := c.Param("crypto")
		price, err := getCryptoPrice(crypto)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, gin.H{"price": price})
	})

	router.Run(":8080")
}
