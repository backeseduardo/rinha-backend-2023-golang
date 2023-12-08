package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

var (
	dbname     = os.Getenv("DB_NAME")
	dbuser     = os.Getenv("DB_USER")
	dbpassword = os.Getenv("DB_PASSWORD")
	dbhost     = os.Getenv("DB_HOST")
)

var DB *sql.DB

func Connect() error {
	uri := fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable", dbuser, dbname, dbpassword, dbhost)

	var err error
	DB, err = sql.Open("postgres", uri)

	return err
}

func Close() {
	if DB != nil {
		DB.Close()
	}
}
