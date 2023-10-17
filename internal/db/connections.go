// db/connection.go
package db

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq" // Import the PostgreSQL driver
)

// Define a global variable for your application configuration.
var AppConfig struct {
	DBHost        string
	DBPort        string
	DBName        string
	DBUser        string
	DBPassword    string
	ServerPort    string
	JWTSecret     string
	JWTExpiration string
	DebugMode     string
	NodeEnv       string
}
var db *sql.DB

// InitDB initializes the database connection.
func InitDB() {
	connStr := fmt.Sprintf(
		"host=%s port=%s user=%s password=%s dbname=%s sslmode=disable",
		AppConfig.DBHost, AppConfig.DBPort, AppConfig.DBUser, AppConfig.DBPassword, AppConfig.DBName,
	)

	var err error
	db, err = sql.Open("postgres", connStr)
	if err != nil {
		log.Fatalf("Error connecting to the database: %v", err)
	}

	// Check the database connection
	err = db.Ping()
	if err != nil {
		log.Fatalf("Error pinging the database: %v", err)
	}
}

// GetDB returns the database connection.
func GetDB() *sql.DB {
	return db
}

// CreateTables creates tables for all models
func CreateTables(db *sql.DB) error {
	sqlStatements := []string{
		`CREATE TABLE IF NOT EXISTS users (
			id UUID PRIMARY KEY,
			first_name VARCHAR(255),
			last_name VARCHAR(255),
			gender VARCHAR(10),
			date_of_birth TIMESTAMP,
			email VARCHAR(255) UNIQUE,
			password VARCHAR(255),
			role VARCHAR(20)
		);`,
		`CREATE TABLE IF NOT EXISTS addresses (
			id UUID PRIMARY KEY,
			user_id UUID,
			street_address VARCHAR(255),
			city VARCHAR(255),
			state VARCHAR(255),
			postal_code VARCHAR(10),
			address_type VARCHAR(20)
		);`,
		`CREATE TABLE IF NOT EXISTS nominees (
			id UUID PRIMARY KEY,
			user_id UUID,
			name VARCHAR(255),
			relationship VARCHAR(20)
		);`,
		`CREATE TABLE IF NOT EXISTS transactions (
			id UUID PRIMARY KEY,
			account_id UUID,
			transaction_type VARCHAR(10),
			amount FLOAT,
			transaction_date TIMESTAMP
		);`,
		`CREATE TABLE IF NOT EXISTS accounts (
			id UUID PRIMARY KEY,
			user_id UUID,
			account_type VARCHAR(20),
			balance FLOAT
		);`,
	}

	for _, sqlStatement := range sqlStatements {
		_, err := db.Exec(sqlStatement)
		if err != nil {
			return err
		}
	}

	return nil
}
