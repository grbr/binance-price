package main

import (
	"fmt"
	"time"

	"github.com/grbr/binance-price/binance"
	"github.com/grbr/binance-price/util"
)

var bitcoinPrice *Price = nil
var fetchError error

func fetchBitcoinPrice() (price *Price, err error) {
	binanceResp, err := binance.SymbolOrderBookTicker(binance.BTCUSDT)
	if err != nil {
		return
	}
	priceAtBinance, err := NewFromSymbolOrderBookTickerR(binanceResp)
	if err != nil {
		return
	}
	price = util.Ptr(ApplyCommission(priceAtBinance))
	return
}

func FetchBitcoinPriceEvery(millis int64) {
	util.SetInterval(func() {
		p, err := fetchBitcoinPrice()
		if err != nil {
			fmt.Println(err)
			if p == nil {
				fetchError = err
			}
			return
		}
		bitcoinPrice = p
	}, time.Duration(millis)*time.Millisecond, true)
}

func GetCachedBitcoinPrice() (price Price, err error) {
	if bitcoinPrice == nil && fetchError != nil {
		return price, fetchError
	}
	return *bitcoinPrice, nil
}
