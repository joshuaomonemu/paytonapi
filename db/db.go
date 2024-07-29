package db

import (
	"database/sql"
	"fmt"
	"net/url"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

type Transaction struct {
	IconUrl string `json:"icon_url"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Amount  string `json:"amount"`
	Status  string `json:"status"`
}

func getDBConnection() (*sql.DB, error) {
	fixieURL, err := url.Parse(os.Getenv("FIXIE_URL"))
	if err != nil {
		return nil, err
	}

	user := "hhhhco_testuser"
	password := "mylovefordogs"
	host := fixieURL.Host
	database := "hhhhco_test"

	dsn := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true", user, password, host, database)
	return sql.Open("mysql", dsn)
}

func GetTrans() string {
	db, err := getDBConnection()
	if err != nil {
		return err.Error()
	}
	defer db.Close()

	rows, err := db.Query("SELECT icon_url, title, date, time, amount, status FROM transactions")
	if err != nil {
		return err.Error()
	}
	defer rows.Close()

	transactions := []Transaction{}
	for rows.Next() {
		var t Transaction
		if err := rows.Scan(&t.IconUrl, &t.Title, &t.Date, &t.Time, &t.Amount, &t.Status); err != nil {
			return err.Error()
		}
		transactions = append(transactions, t)
	}

	if err := rows.Err(); err != nil {
		return err.Error()
	}

	return fmt.Sprintf("%+v\n", transactions)
}
