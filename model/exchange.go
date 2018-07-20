package model

type Exchange struct {
	ID           int64   `json:"id,omitempty"`
	ExchangeDate string  `json:"exchange_date"`
	CurrencyFrom string  `json:"currency_from"`
	CurrencyTo   string  `json:"currency_to"`
	Rate         float64 `json:"rates"`
	CreatedByIp  string  `json:"created_by_ip"`
	UpdatedByIp  string  `json:"updated_by_ip"`
	CreatedAt    string  `json:"created_at,omitempty"`
	UpdatedAt    string  `json:"updated_at,omitempty"`
}

func NewExchangeModel() Exchange {
	return Exchange{}
}
