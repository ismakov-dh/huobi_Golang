package linearswap

type HistoryOrder struct {
	ContractType    string  `json:"contract_type"`
	Pair            string  `json:"pair"`
	BusinessType    string  `json:"business_type"`
	OrderId         int64   `json:"order_id"`
	ContractCode    string  `json:"contract_code"`
	Symbol          string  `json:"symbol"`
	LeverRate       int     `json:"lever_rate"`
	Direction       string  `json:"direction"`
	Offset          string  `json:"offset"`
	Volume          float64 `json:"volume"`
	Price           float64 `json:"price"`
	CreateDate      int64   `json:"create_date"`
	OrderSource     string  `json:"order_source"`
	OrderPriceType  int     `json:"order_price_type"`
	OrderType       int     `json:"order_type"`
	MarginFrozen    float64 `json:"margin_frozen"`
	Profit          float64 `json:"profit"`
	TradeVolume     float64 `json:"trade_volume"`
	TradeTurnover   float64 `json:"trade_turnover"`
	Fee             float64 `json:"fee"`
	TradeAvgPrice   float64 `json:"trade_avg_price"`
	Status          int     `json:"status"`
	OrderIdStr      string  `json:"order_id_str"`
	FeeAsset        string  `json:"fee_asset"`
	LiquidationType string  `json:"liquidation_type"`
	MarginAsset     string  `json:"margin_asset"`
	MarginMode      string  `json:"margin_mode"`
	MarginAccount   string  `json:"margin_account"`
	UpdateTime      int64   `json:"update_time"`
	IsTpsl          int     `json:"is_tpsl"`
	RealProfit      int     `json:"real_profit"`
	ReduceOnly      int     `json:"reduce_only"`
}
