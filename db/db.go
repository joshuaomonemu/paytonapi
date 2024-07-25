package db

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
)

func InsertReq() {
	db, err := connectToMySQL()
	if err != nil {
		panic(err)
	}
	defer db.Close() // Close the database connection when the function exits

	// Sample query to select data from a table
	rows, err := db.Query("SELECT * FROM your_table_name")
	if err != nil {
		panic(err)
	}
	defer rows.Close() // Close the result set

	// Scan the results
	for rows.Next() {
		var id int
		var name string
		err := rows.Scan(&id, &name)
		if err != nil {
			panic(err)
		}
		fmt.Println("ID:", id, "Name:", name)
	}

	if err := rows.Err(); err != nil {
		panic(err)
	}
}

func connectToMySQL() (*sql.DB, error) {
	// Replace with your actual credentials
	dbUser := "hhhhco_testuser"
	dbPassword := "mylovefordogs"
	dbName := "hhhhco_test"
	dbHost := "54.38.50.173" // This could be "localhost" or the cPanel server's hostname/IP
	dbPort := "3306"         // Standard MySQL port

	connectionString := dbUser + ":" + dbPassword + "@tcp(" + dbHost + ":" + dbPort + ")/" + dbName

	db, err := sql.Open("mysql", connectionString)
	if err != nil {
		return nil, err
	}

	if err := db.Ping(); err != nil {
		return nil, err
	}

	return db, nil
}
