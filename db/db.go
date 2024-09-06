package db

import (
	"database/sql"
	"encoding/json"
	"errors"
	"strconv"

	_ "github.com/go-sql-driver/mysql"
)

// Struct to hold table data
type User struct {
	ID     int
	FName  string
	LName  string
	Email  string
	Wallet string
}
type Transaction struct {
	IconUrl string `json:"icon_url"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Amount  string `json:"amount"`
	Status  string `json:"status"`
	User    string `json:"user"`
}

type TransactionPayload struct {
	IconUrl string `json:"icon_url"`
	Title   string `json:"title"`
	Date    string `json:"date"`
	Time    string `json:"time"`
	Amount  string `json:"amount"`
	Status  string `json:"status"`
}

func Conn() (*sql.DB, error) {
	// Database connection string
	// Format: username:password@tcp(localhost:3306)/dbname
	dsn := "spades:mylovefordogs@tcp(payton.c1ws4goq8w6a.eu-north-1.rds.amazonaws.com:3306)/ineracsi_payment_app"
	//dsn := "root:@tcp(127.0.0.1:3306)/test"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//defer db.Close()

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		return nil, err
	}
	return db, nil
}

func GetTransactions(id string) ([]byte, error) {
	db, err1 := Conn()
	if err1 != nil {
		return nil, err1
	}
	// Query to fetch data from the table
	rows, err := db.Query("SELECT iconurl, title, date, time, amount, status FROM transactions WHERE user = ?", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	// Slice to hold the results
	var transactions []TransactionPayload

	// Iterate over the rows
	for rows.Next() {
		var transaction TransactionPayload
		err := rows.Scan(&transaction.IconUrl, &transaction.Title, &transaction.Date, &transaction.Time, &transaction.Amount, &transaction.Status)
		if err != nil {
			return nil, err
		}
		transactions = append(transactions, transaction)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// Print the results
	//for _, transaction := range transactions {
	//	fmt.Printf("Iconurl: %s, Title: %s, Date: %s, Time: %s, Amount: %s, Status: %s\n", transaction.IconUrl, transaction.Title, transaction.Date, transaction.Time, transaction.Amount, transaction.Status)
	//}
	bs, _ := json.Marshal(transactions)
	return bs, nil

}

func SetTransaction(transaction *Transaction) error {
	db, err1 := Conn()
	if err1 != nil {
		return err1
	}
	query := `INSERT INTO transactions (iconurl, title, date, time, amount, status, user)
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, transaction.IconUrl, transaction.Title, transaction.Date, transaction.Time, transaction.Amount, transaction.Status, transaction.User)
	return err
}

func GetUser() ([]User, error) {
	db, err := Conn()
	if err != nil {
		return nil, err
	}
	// Query to fetch data from the table
	rows, err := db.Query("SELECT id, fname, lname, email, wallet FROM users")
	if err != nil {
		return nil, err
	}
	//defer rows.Close()

	// Slice to hold the results
	var users []User

	// Iterate over the rows
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FName, &user.LName, &user.Email, &user.Wallet)
		if err != nil {
			return nil, err
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		return nil, err
	}

	// Print the results
	// for _, user := range users {
	// 	fmt.Printf("ID: %d, FName: %s, LName: %s, Email: %s\n", user.ID, user.FName, user.LName, user.Email)
	// }
	return users, nil
}

func CheckBalance(amount, email string) (string, error) {
	db, err := Conn()
	if err != nil {
		return "", err
	}
	// Convert amount to float64
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}

	var wallet float64
	query := "SELECT wallet FROM users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no user found with the provided email")
		}
		return "", err
	}

	if amt > wallet {
		return "", errors.New("error")
	}

	return "okay", nil
}

func LoadWallet(email string) (float64, error) {
	db, err := Conn()
	if err != nil {
		return 0, err
	}

	var wallet float64
	query := "SELECT wallet FROM users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return 0, errors.New("no user found with the provided email")
		}
		return 0, err
	}
	return wallet, nil
}

func UpdateBalance(email, amount string) error {
	db, err1 := Conn()
	if err1 != nil {
		return err1
	}
	query := `UPDATE users SET wallet = ? WHERE email = ?`
	_, err := db.Exec(query, amount, email)
	return err
}

func WalletTrans(amount, email string) (string, error) {
	db, err := Conn()
	if err != nil {
		return "", err
	}
	// Convert amount to float64
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return "", err
	}

	var wallet float64
	query := "SELECT wallet FROM users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return "", errors.New("no user found with the provided email")
		}
		return "", err
	}
	new_bal := wallet - amt
	query1 := `UPDATE users SET wallet = ? WHERE email = ?`

	_, err = db.Exec(query1, new_bal, email)
	if err != nil {
		return "", err
	}

	return "okay", nil
}
