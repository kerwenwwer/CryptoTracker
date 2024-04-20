package main

import (
	"log"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"

	"github.com/kerwenwwer/CryptoTracker/pkg/api"
)

func main() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Error loading .env file")
	}

	router := gin.Default()

	// Set trusted proxies
	trustedProxies := []string{} // Now is empy, but we can add trusted proxies
	if err := router.SetTrustedProxies(trustedProxies); err != nil {
		log.Fatal("Failed to set trusted proxies")
	}

	// Setup CORS if necessary
	router.Use(cors.Default())

	// Serve static files from Vue.js build directory
	router.Static("/app", "./frontend/dist")

	router.GET("/price/:crypto", func(c *gin.Context) {
		crypto := c.Param("crypto")
		price, err := api.FetchCryptoPrice(crypto)
		if err != nil {
			c.JSON(500, gin.H{"error": err.Error()})
			return
		}
		c.JSON(200, gin.H{"price": price})
	})

	router.Run(":8080")
}
