package trader

import (
	"log"

	"github.com/kerwenwwer/CryptoTracker/pkg/api"
)

func ExecuteTrade(crypto string) {
	price, err := api.FetchCryptoPrice(crypto)
	if err != nil {
		log.Printf("Error fetching price for %s: %v", crypto, err)
		return
	}
	log.Printf("Current price of %s is %f USD", crypto, price)
	// Add trading logic here: when to buy/sell based on price movements or other indicators
}
