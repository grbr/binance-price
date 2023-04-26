package binance

import (
	"bufio"
	"encoding/json"
	"fmt"
	"net/http"
)

const BINANCE_API = "https://api.binance.com"

type Symbol = string

const (
	BTCUSDT Symbol = "BTCUSDT" // Bitcoin
)

type SymbolOrderBookTickerR = struct {
	Symbol   string `json:"symbol"`
	BidPrice string `json:"bidPrice"`
	BidQty   string `json:"bidQty"`
	AskPrice string `json:"askPrice"`
	AskQty   string `json:"askQty"`
}

func SymbolOrderBookTicker(s Symbol) (result SymbolOrderBookTickerR, err error) {
	data, err := httpGetJson(BINANCE_API + "/api/v3/ticker/bookTicker?symbol=" + s)
	if err != nil {
		fmt.Println(err)
		return result, err
	}
	json.Unmarshal([]byte(data), &result)
	return result, err
}

func httpGetJson(url string) (data string, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		data += scanner.Text()
	}
	err = scanner.Err()

	return data, err
}
