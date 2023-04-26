package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/grbr/binance-price/util"
	"github.com/shopspring/decimal"
)

func priceBitcoinBinance(c *gin.Context) {
	price, err := GetCachedBitcoinPrice()
	if err != nil {
		fmt.Println(err)
		c.JSON(http.StatusServiceUnavailable, "Service unavailable")
	} else {
		c.JSON(http.StatusOK, price.toDTO())
	}
}

func setupRouter(r *gin.Engine) {
	r.GET("/", priceBitcoinBinance)
}

func checkConfigAndSetDefaults(config *util.Config) {
	if config.PORT == 0 {
		config.PORT = 3000
	}
	if config.PORT < 2 || config.PORT > 65536 {
		log.Fatal("[config] PORT must be between 2 and 65536")
	}
	if config.UPDATE_INTERVAL_MILLIS == 0 {
		config.UPDATE_INTERVAL_MILLIS = 10000
	}
	if config.UPDATE_INTERVAL_MILLIS < 0 {
		log.Fatal("[config] UPDATE_INTERVAL_MILLIS must be a positive number")
	}
	if config.SERVICE_COMMISSION_PERCENT <= 0 {
		log.Fatal("[config] SERVICE_COMMISSION_PERCENT must be a positive number")
	}
}

func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal("cannot load config: ", err)
	}
	checkConfigAndSetDefaults(&config)

	SetCommissionPercent(decimal.NewFromFloat(config.SERVICE_COMMISSION_PERCENT))
	FetchBitcoinPriceEvery(int64(config.UPDATE_INTERVAL_MILLIS))

	r := gin.Default()
	setupRouter(r)
	err = r.Run(fmt.Sprintf(":%d", config.PORT))
	if err != nil {
		log.Fatalf("gin Run error: %s", err)
	}
}
