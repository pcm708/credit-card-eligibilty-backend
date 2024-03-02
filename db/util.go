package db

import (
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
	. "github.com/honestbank/tech-assignment-backend-engineer/constants"
	"github.com/joho/godotenv"
	"log"
	"os"
)

var db *sql.DB

func ConnectToDB() {
	var err error
	var dbURL string

	// Load environment variables from .env file
	err = godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	if os.Getenv("CLOUD") == "true" {
		dbURL = "root:root@tcp(" + CLOUD_DB_URL + ":" + DB_PORT + ")/number"
	} else {
		dbURL = "root:root@tcp(db:" + DB_PORT + ")/number"
	}

	db, err = sql.Open("mysql", dbURL)
	if err != nil {
		log.Fatal(err)
	}

	if err = db.Ping(); err != nil {
		log.Fatal(err)
	}

	// Check if the table exists and create it if it doesn't
	err = createTableIfNotExists()
	if err != nil {
		log.Fatal(err)
	}
}

func createTableIfNotExists() error {
	query := `CREATE TABLE IF NOT EXISTS phone_numbers (
  number VARCHAR(255) PRIMARY KEY
 )`
	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("error creating table: %v", err.Error())
	}

	return nil
}

func StoreNewNumber(phoneNumber string) error {
	query := `INSERT INTO phone_numbers (number) VALUES (?)`
	_, err := db.Exec(query, phoneNumber)
	if err != nil {
		return err
	}
	return nil
}

func CheckIfNumberPresent(number string) (bool, error) {
	var exists bool
	query := `SELECT exists (SELECT 1 FROM phone_numbers WHERE number=?)`
	err := db.QueryRow(query, number).Scan(&exists)
	if err != nil && err != sql.ErrNoRows {
		return false, err
	}
	return exists, nil
}
