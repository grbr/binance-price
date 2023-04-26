package main

import (
	"github.com/grbr/binance-price/binance"
	"github.com/shopspring/decimal"
)

type Price struct {
	Ask decimal.Decimal
	Bid decimal.Decimal
}

type PriceDto struct {
	Ask string `json:"ask"`
	Bid string `json:"bid"`
	Mid string `json:"mid"`
}

func (p Price) toDTO() PriceDto {
	return PriceDto{
		p.Ask.StringFixed(8),
		p.Bid.StringFixed(8),
		p.Ask.Add(p.Bid).Div(decimal.NewFromInt(2)).StringFixed(8),
	}
}

func NewFromSymbolOrderBookTickerR(s binance.SymbolOrderBookTickerR) (price Price, err error) {
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
