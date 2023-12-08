package database

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/joho/godotenv/autoload"
	_ "github.com/lib/pq"
)

func Init() (*sql.DB, error) {
	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		uri        = fmt.Sprintf("user=%s dbname=%s password=%s host=%s port=5432 sslmode=disable", dbuser, dbname, dbpassword, dbhost)
	)

	db, err := sql.Open("postgres", uri)
	if err != nil {
		return nil, err
	}

	return db, nil
}
