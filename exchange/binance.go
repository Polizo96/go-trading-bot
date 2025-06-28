package exchange

import (
	"encoding/json"
	"fmt"
	"go-trading-bot/config"
	"net/http"
)

func GetPrices(symbol string, limit int) ([]float64, error) {
	url := fmt.Sprintf("%s/api/v3/klines?symbol=%s&interval=1m&limit=%d", config.BinanceBaseURL, symbol, limit)

	resp, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var raw [][]interface{}
	if err := json.NewDecoder(resp.Body).Decode(&raw); err != nil {
		return nil, err
	}

	var closes []float64
	for _, item := range raw {
		closePriceStr := item[4].(string)
		var closePrice float64
		fmt.Sscanf(closePriceStr, "%f", &closePrice)
		closes = append(closes, closePrice)
	}
	return closes, nil
}
