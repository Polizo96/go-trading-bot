package types

type TradeSignal int

const (
	SignalHold TradeSignal = iota
	SignalBuy
	SignalSell
)

func (s TradeSignal) String() string {
	return [...]string{"HOLD", "BUY", "SELL"}[s]
}
