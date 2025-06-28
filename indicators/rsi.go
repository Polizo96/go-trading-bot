package indicators

func CalculateRSI(closes []float64) float64 {
	if len(closes) < 15 {
		return 50.0
	}

	var gains, losses float64
	for i := 1; i < len(closes); i++ {
		diff := closes[i] - closes[i-1]
		if diff > 0 {
			gains += diff
		} else {
			losses -= diff
		}
	}
	avgGain := gains / float64(len(closes)-1)
	avgLoss := losses / float64(len(closes)-1)

	if avgLoss == 0 {
		return 100.0
	}
	rs := avgGain / avgLoss
	return 100.0 - (100.0 / (1 + rs))
}
