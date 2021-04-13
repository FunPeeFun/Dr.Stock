package buffer

import "sync"

type Config struct {
	TotalAmount float32 `json:"total_amount"`
	SetTotalPercent float32 `json:"set_total_percent"`
	FixedInvestWeek int32 `json:"fixed_invest_week"`
	Stocks map[string]Stock
}

type Stock struct {
	Name string
	Role string
	Amount float32
	SetPercent float32 `json:"set_percent"`
}

var StockConfig Config
var StockConfigLock sync.Mutex