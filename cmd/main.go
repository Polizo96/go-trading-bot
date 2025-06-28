package main

import (
	"go-trading-bot/tradeengine"
)

func main() {
	engine := tradeengine.NewEngine()
	engine.Start()
}
