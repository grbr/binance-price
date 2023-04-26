package main

import (
	"github.com/shopspring/decimal"
)

var commission decimal.Decimal
var rate decimal.Decimal
var commissionSet bool = false

func SetCommissionPercent(percent decimal.Decimal) {
	one := decimal.NewFromInt(1)
	ahundred := decimal.NewFromInt(100)
	rate = one.Add(percent.Div(ahundred))
	commission = percent
	commissionSet = true
}

func ApplyCommission(p Price) Price {
	if !commissionSet {
		panic("Must set commission")
	}
	return Price{p.Ask.Mul(rate), p.Bid.Mul(rate)}
}
