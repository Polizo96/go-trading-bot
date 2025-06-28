package strategy

import (
	"go-trading-bot/config"
	"go-trading-bot/indicators"
	"go-trading-bot/types"
)

func EvaluateRSIStrategy(prices []float64) string {
	rsi := indicators.CalculateRSI(prices)

	if rsi < config.BuyThreshold {
		return types.SignalBuy.String()
	}
	if rsi > config.SellThreshold {
		return types.SignalSell.String()
	}
	return types.SignalHold.String()
}
