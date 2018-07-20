package model

import "github.com/teguhmaulana25/shopee_tech/config/database"

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

func (m Exchange) All() []Exchange {
	var allData []Exchange
	rows, err := database.Db.Query("SELECT * FROM sp_exchange_rates")
	database.Check(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&m.ID,
			&m.ExchangeDate,
			&m.CurrencyFrom,
			&m.CurrencyTo,
			&m.Rate,
			&m.CreatedByIp,
			&m.UpdatedByIp,
			&m.CreatedAt,
			&m.UpdatedAt,
		)
		database.Check(err)
		allData = append(allData, m)
	}

	return allData
}

func (m Exchange) Store() (database.DBUpdate, error) {
	// Create prepared statement.
	query := "INSERT INTO sp_exchange_rates (exchange_date, currency_from, currency_to, rate, created_by_ip, updated_by_ip, created_at, updated_at) VALUES(?, ?, ?, ?, ?, ?, ?, ?)"
	stmt, err := database.Db.Prepare(query)
	database.Check(err)

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(
		&m.ExchangeDate,
		&m.CurrencyFrom,
		&m.CurrencyTo,
		&m.Rate,
		&m.CreatedByIp,
		&m.UpdatedByIp,
		&m.CreatedAt,
		&m.UpdatedAt,
	)
	database.Check(err)
	lastID, err := res.LastInsertId()
	database.Check(err)
	rowCnt, err := res.RowsAffected()
	database.Check(err)

	// Populate DBUpdate struct with last Id and num rows affected.
	database.UpdateResult.ID = int(lastID)
	database.UpdateResult.Affected = rowCnt

	return database.UpdateResult, nil
}
