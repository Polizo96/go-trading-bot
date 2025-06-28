package tradeengine

import (
	"fmt"
	"go-trading-bot/config"
	"go-trading-bot/exchange"
	"go-trading-bot/strategy"
	"time"
)

type Engine struct{}

func NewEngine() *Engine {
	return &Engine{}
}

func (e *Engine) Start() {
	for {
		prices, err := exchange.GetPrices(config.Symbol, config.RSIPeriod+1)
		if err != nil {
			fmt.Println("Error fetching prices:", err)
			continue
		}

		action := strategy.EvaluateRSIStrategy(prices)
		fmt.Printf("[%s] Action: %s | Last price: %.2f\n", time.Now().Format("15:04:05"), action, prices[len(prices)-1])

		// TODO: Record action to blockchain or simulate trade

		time.Sleep(60 * time.Second)
	}
}
