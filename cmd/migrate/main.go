package main

import (
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/joho/godotenv/autoload"
)

func main() {
	var (
		dbname     = os.Getenv("DB_NAME")
		dbuser     = os.Getenv("DB_USER")
		dbpassword = os.Getenv("DB_PASSWORD")
		dbhost     = os.Getenv("DB_HOST")
		uri        = fmt.Sprintf("postgres://%s:%s@%s:5432/%s?sslmode=disable", dbuser, dbpassword, dbhost, dbname)
	)

	m, err := migrate.New("file://database/migrations", uri)
	if err != nil {
		log.Fatalln(err)
	}

	defer m.Close()

	arg := os.Args[1]

	if arg == "up" {
		m.Up()
		return
	}

	if arg == "down" {
		m.Down()
		return
	}
}
