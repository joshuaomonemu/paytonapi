package db

import (
	structs "app/struct"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strconv"
	"time"

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
type User2 struct {
	Fullame string
	Phone   string
	Email   string
	Wallet  string
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
	dsn := "root:mypassword@tcp(localhost:3306)/paytondb"
	//dsn := "root:@tcp(127.0.0.1:3306)/test"

	// Open a connection to the database
	db, err := sql.Open("mysql", dsn)
	if err != nil {
		return nil, err
	}
	//defer db.Close()
	db.SetConnMaxLifetime(time.Minute * 3)

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

func TransApprove(id string) ([]byte, error) {
	db, err1 := Conn()
	if err1 != nil {
		return nil, err1
	}
	// Query to fetch data from the table
	rows, err := db.Query("SELECT iconurl, title, date, time, amount, status FROM transactions WHERE user = ? AND status = ?", id, "Approved")
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

func GetWalletTransactions(id string) ([]byte, error) {
	db, err1 := Conn()
	if err1 != nil {
		return nil, err1
	}
	// Query to fetch data from the table
	rows, err := db.Query("SELECT iconurl, title, date, time, amount, status FROM wallets WHERE user = ?", id)
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

func SetWallets(transaction *Transaction) error {
	db, err1 := Conn()
	if err1 != nil {
		return err1
	}
	query := `INSERT INTO wallets (iconurl, title, date, time, amount, status, user)
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

func SetUser(user *structs.UserData) error {
	db, err1 := Conn()
	if err1 != nil {
		return err1
	}
	query := `INSERT INTO user1 (email, fullname, phone, password)
              VALUES (?, ?, ?, ?)`
	_, err := db.Exec(query, user.Email, user.Fullname, user.Phone, user.Password)

	return err
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

func UpdateWallet(email, amount string) error {
	db, err1 := Conn()
	if err1 != nil {
		return err1
	}

	// Convert amount to float64
	amt, err := strconv.ParseFloat(amount, 64)
	if err != nil {
		return err
	}

	var wallet float64
	query := "SELECT wallet FROM users WHERE email = ?"
	err = db.QueryRow(query, email).Scan(&wallet)
	if err != nil {
		if err == sql.ErrNoRows {
			return errors.New("no user found with the provided email")
		}
		return err
	}

	new_balance := amt + wallet

	query1 := `UPDATE users SET wallet = ? WHERE email = ?`
	_, err = db.Exec(query1, new_balance, email)

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

func EmailExists(email string) (bool, error) {
	db, _ := Conn()

	const query = "SELECT 1 FROM user1 WHERE email = ?"
	row := db.QueryRow(query, email)
	var exists bool
	err := row.Scan(&exists)
	if err != nil {
		if err == sql.ErrNoRows {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

func StoreOTP(email string, otp string) error {
	db, _ := Conn()
	// Store the OTP in the database with a timestamp
	expires := time.Now().Add(10 * time.Minute)
	_, err := db.Exec("INSERT INTO otp (email, pin, expires) VALUES (?, ?, ?)", email, otp, expires)
	return err
}

func GetOTP(email string) (string, error) {
	db, _ := Conn()
	// Retrieve the OTP from the database
	var otp string
	var expires []byte // Change the type to []byte to hold the raw data
	err := db.QueryRow("SELECT pin, expires FROM otp WHERE email = ?", email).Scan(&otp, &expires)
	if err != nil {
		return "", err
	}

	// Convert []byte to string and then parse as time.Time
	parsedTime, err := time.Parse("2006-01-02 15:04:05", string(expires))
	if err != nil {
		return "", fmt.Errorf("error parsing expiration time: %v", err)
	}

	// Check if the OTP has expired
	if time.Now().After(parsedTime) {
		return "", fmt.Errorf("OTP has expired")
	}
	return otp, nil
}

func LoginUser(email, password string) (bool, error) {
	db, _ := Conn()
	// Retrieve the OTP from the database

	var exists bool

	// Prepare the SQL query
	query := "SELECT EXISTS(SELECT 1 FROM user1 WHERE email = ? AND password = ?)"

	// Execute the query
	err := db.QueryRow(query, email, password).Scan(&exists)
	if err != nil {
		return false, err
	}

	// Return the existence of the user with the given email and password
	return exists, nil
}

func GetUserbyEmail(email string) (User2, error) {
	db, _ := Conn()
	var user User2

	// Prepare the SQL query
	query := "SELECT fullname, phone, wallet, email FROM user1 WHERE email = ?"

	// Execute the query
	err := db.QueryRow(query, email).Scan(&user.Fullame, &user.Phone, &user.Wallet, &user.Email)
	if err != nil {
		if err == sql.ErrNoRows {
			return User2{}, err
		}
		return User2{}, err
	}

	return user, nil
}
