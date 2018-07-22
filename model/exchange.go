package model

import (
	"github.com/teguhmaulana25/shopee_tech/config/database"
)

type Exchange struct {
	ID           int64   `json:"id,omitempty"`
	ExchangeDate string  `json:"exchange_date"`
	CurrencyFrom string  `json:"currency_from"`
	CurrencyTo   string  `json:"currency_to"`
	Rate         float64 `json:"rate"`
	AverageRate  string  `json:"average_rate"`
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

func (Exchange) Delete(from string, to string) database.DBUpdate {
	// Create prepared statement.
	stmt, err := database.Db.Prepare("DELETE FROM sp_exchange_rates where currency_from=? and currency_to=?")
	database.Check(err)

	// Execute the prepared statement and retrieve the results.
	res, err := stmt.Exec(
		from,
		to,
	)

	database.Check(err)
	rowCnt, err := res.RowsAffected()
	database.Check(err)

	// Populate DBUpdate struct with last id and num rows affected.
	database.UpdateResult.ID = 0
	database.UpdateResult.Affected = rowCnt

	return database.UpdateResult
}

func (m Exchange) Tracked(current_date string, last_date string) []Exchange {
	var allData []Exchange
	query := "SELECT avg(rate) as average_rate, exchange_date, currency_from, currency_to, rate FROM sp_exchange_rates WHERE exchange_date>='" + last_date + "' and exchange_date<='" + current_date + "' GROUP BY currency_from, currency_to"
	rows, err := database.Db.Query(query)
	database.Check(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&m.AverageRate,
			&m.ExchangeDate,
			&m.CurrencyFrom,
			&m.CurrencyTo,
			&m.Rate,
		)
		database.Check(err)
		allData = append(allData, m)
	}
	return allData
}

func (m Exchange) Find(currencyFrom string, currencyTo string) []Exchange {
	var allData []Exchange

	query := "SELECT id, exchange_date, currency_from, currency_to, rate FROM sp_exchange_rates WHERE currency_from='" + currencyFrom + "' and currency_to='" + currencyTo + "' ORDER BY id DESC"
	rows, err := database.Db.Query(query)
	database.Check(err)
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(
			&m.ID,
			&m.ExchangeDate,
			&m.CurrencyFrom,
			&m.CurrencyTo,
			&m.Rate,
		)
		database.Check(err)
		allData = append(allData, m)
	}

	return allData
}
