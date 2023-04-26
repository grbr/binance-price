package main

import (
	"time"

	"github.com/grbr/binance-price/binance"
	"github.com/grbr/binance-price/utils"
	"github.com/shopspring/decimal"
)

var cachedBitcoinPrice Price

func binancePrice(s binance.SymbolOrderBookTickerR) (price Price, err error) {
	ask, err := decimal.NewFromString(s.AskPrice)
	if err != nil {
		return
	}
	bid, err := decimal.NewFromString(s.BidPrice)
	if err != nil {
		return
	}
	return Price{ask, bid}, err
}

func updateBitcoinPrice() (price Price, err error) {
	binanceResp, err := binance.SymbolOrderBookTicker(binance.BTCUSDT)
	if err != nil {
		return
	}
	priceAtBinance, err := binancePrice(binanceResp)
	if err != nil {
		return
	}
	price = ApplyCommission(priceAtBinance)
	cachedBitcoinPrice = price
	return
}

func CacheBitcoinPriceEvery(millis int64) {
	utils.SetInterval(func() {
		updateBitcoinPrice()
	}, time.Duration(millis)*time.Millisecond, true)
}

func GetCachedBitcoinPrice() (Price, error) {
	return cachedBitcoinPrice, nil
}
