package api

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"os"
)

type CryptoData struct {
	Price float64 `json:"rate"`
}

func FetchCryptoPrice(crypto string) (float64, error) {
	apiKey := os.Getenv("COINAPI_KEY")
	url := fmt.Sprintf("https://rest.coinapi.io/v1/exchangerate/%s/USD", crypto)
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return 0, err
	}

	req.Header.Add("X-CoinAPI-Key", apiKey)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return 0, err
	}

	var data CryptoData
	if err = json.Unmarshal(body, &data); err != nil {
		return 0, err
	}

	return data.Price, nil
}
