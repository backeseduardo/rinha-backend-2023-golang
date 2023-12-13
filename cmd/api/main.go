package main

import (
	"github.com/backeseduardo/rinha-backend-2023-golang/internal/database"
	"github.com/backeseduardo/rinha-backend-2023-golang/internal/http"
	"github.com/backeseduardo/rinha-backend-2023-golang/internal/person"
)

func main() {
	database.Connect()
	defer database.Close()

	pRepository := person.NewDBRepository()

	http.NewServer(&http.HttpServerOpts{
		PersonRepository: pRepository,
	})

	/*
		users, err := getUsers(db)
		if err != nil {
			log.Fatalln(err)
		}

		for _, user := range users {
			log.Println(user.Username)
		}
	*/
}
