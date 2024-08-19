package db

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net"
	"net/http"
	"net/url"
	"os"
	"strconv"

	"github.com/go-sql-driver/mysql"
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

func ProxyConn(fixieUrl string, username string, password string, dbName string) (*sql.DB, error) {
	// Parse the Fixie URL
	fixieProxyUrl, err := url.Parse(fixieUrl)
	if err != nil {
		return nil, err
	}

	// Create a new HTTP transport with the Fixie proxy
	httpTransport := &http.Transport{
		Proxy: http.ProxyURL(fixieProxyUrl),
	}

	// Register a custom dialer for the MySQL driver
	mysql.RegisterDialContext("mysql-proxy", func(ctx context.Context, addr string) (net.Conn, error) {
		// Create a new proxy connection to the remote database
		return httpTransport.DialContext(ctx, "tcp", addr)
	})

	// Create a new database connection using the proxy dialer
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(mysql-proxy)/%s", username, password, dbName))
	if err != nil {
		return nil, err
	}

	return db, nil
}

func Conn() *sql.DB {
	// Database connection string
	// Format: username:password@tcp(localhost:3306)/dbname
	//dsn := "ineracsi_baker:Goodmorning11.@tcp(54.38.50.173:2083)/ineracsi_payment_app"
	//dsn := "root:@tcp(127.0.0.1:3306)/test"

	// Open a connection to the database
	// db, err := sql.Open("mysql", dsn)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	//defer db.Close()
	// Get the Fixie URL from the Heroku environment variable
	fixieUrl := os.Getenv("http://fixie:UbuzTDFXKE2kciz@velodrome.usefixie.com:80")
	username := "ineracsi_baker"
	password := "Goodmorning11."
	dbName := "ineracsi_payment_app"

	// Open the database connection
	db, err := ProxyConn(fixieUrl, username, password, dbName)
	if err != nil {
		fmt.Println(err)

	}

	// Ping the database to verify connection
	err = db.Ping()
	if err != nil {
		log.Fatal(err)
	}
	return db
}

func GetTransactions(id string) []byte {
	db := Conn()
	// Query to fetch data from the table
	rows, err := db.Query("SELECT iconurl, title, date, time, amount, status FROM transactions WHERE user = ?", id)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to hold the results
	var transactions []TransactionPayload

	// Iterate over the rows
	for rows.Next() {
		var transaction TransactionPayload
		err := rows.Scan(&transaction.IconUrl, &transaction.Title, &transaction.Date, &transaction.Time, &transaction.Amount, &transaction.Status)
		if err != nil {
			log.Fatal(err)
		}
		transactions = append(transactions, transaction)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	//for _, transaction := range transactions {
	//	fmt.Printf("Iconurl: %s, Title: %s, Date: %s, Time: %s, Amount: %s, Status: %s\n", transaction.IconUrl, transaction.Title, transaction.Date, transaction.Time, transaction.Amount, transaction.Status)
	//}
	bs, _ := json.Marshal(transactions)
	return bs

}

func SetTransaction(transaction *Transaction) error {
	db := Conn()
	query := `INSERT INTO transactions (iconurl, title, date, time, amount, status, user) 
              VALUES (?, ?, ?, ?, ?, ?, ?)`
	_, err := db.Exec(query, transaction.IconUrl, transaction.Title, transaction.Date, transaction.Time, transaction.Amount, transaction.Status, transaction.User)
	return err
}

func GetUser() []User {
	db := Conn()
	// Query to fetch data from the table
	rows, err := db.Query("SELECT id, fname, lname, email, wallet FROM users")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// Slice to hold the results
	var users []User

	// Iterate over the rows
	for rows.Next() {
		var user User
		err := rows.Scan(&user.ID, &user.FName, &user.LName, &user.Email, &user.Wallet)
		if err != nil {
			log.Fatal(err)
		}
		users = append(users, user)
	}

	// Check for errors from iterating over rows
	err = rows.Err()
	if err != nil {
		log.Fatal(err)
	}

	// Print the results
	// for _, user := range users {
	// 	fmt.Printf("ID: %d, FName: %s, LName: %s, Email: %s\n", user.ID, user.FName, user.LName, user.Email)
	// }
	return users
}

func CheckBalance(amount, email string) (error, string) {
	db := Conn()
	// Convert amount to float64
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err, ""
	}

	var wallet float64
	query := "SELECT wallet FROM users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no user found with the provided email"), ""
		}
		return err, ""
	}

	if amt > wallet {
		return errors.New("error"), ""
	}

	return nil, "okay"
}

func UpdateBalance(email, amount string) error {
	db := Conn()
	query := `UPDATE users SET wallet = ? WHERE email = ?`
	_, err := db.Exec(query, amount, email)
	return err
}

func WalletTrans(amount, email string) (error, string) {
	db := Conn()
	// Convert amount to float64
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err, ""
	}

	var wallet float64
	query := "SELECT wallet FROM users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no user found with the provided email"), ""
		}
		return err, ""
	}
	new_bal := wallet - amt
	query1 := `UPDATE users SET wallet = ? WHERE email = ?`

	_, err = db.Exec(query1, new_bal, email)
	if err != nil {
		return err, ""
	}

	return nil, "okay"
}
