package main

import "github.com/shopspring/decimal"

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
