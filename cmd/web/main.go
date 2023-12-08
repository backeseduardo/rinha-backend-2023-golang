package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/backeseduardo/rinha-backend-2023-golang"
	"github.com/backeseduardo/rinha-backend-2023-golang/database"
	"github.com/gin-gonic/gin"
)

func main() {
	db, err := database.Init()
	if err != nil {
		log.Fatalln(err)
	}

	defer db.Close()

	HandlePostPessoas := func(c *gin.Context) {
		var pessoaBody rinha.Pessoa

		err := json.NewDecoder(c.Request.Body).Decode(&pessoaBody)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  500,
				"message": err.Error(),
			})
			return
		}

		id, err := rinha.InsertPerson(db, &pessoaBody)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  500,
				"message": err.Error(),
			})
			return
		}

		c.Header("Location", fmt.Sprintf("/pessoas/%d", id))
		c.JSON(http.StatusCreated, gin.H{})
	}

	HandleGetPessoaById := func(c *gin.Context) {
		id := c.Param("id")

		p, err := rinha.GetPerson(db, id)
		if err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{
				"status":  500,
				"message": err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, p)
	}

	HandleGetPessoas := func(c *gin.Context) {
		c.JSON(http.StatusInternalServerError, gin.H{})
	}

	HandleGetContagemPessoas := func(c *gin.Context) {
		c.String(http.StatusOK, "1000")
	}

	r := gin.Default()
	r.POST("/pessoas", HandlePostPessoas)
	r.GET("/pessoas/:id", HandleGetPessoaById)
	r.GET("/pessoas", HandleGetPessoas)
	r.GET("/contagem-pessoas", HandleGetContagemPessoas)
	r.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": c.Param("id"),
		})
	})
	r.Run()

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

type user struct {
	Username string
}

func getUsers(db *sql.DB) ([]user, error) {
	rows, err := db.Query("SELECT username FROM users")
	if err != nil {
		return nil, err
	}

	defer rows.Close()

	var users []user

	for rows.Next() {
		var u user

		err = rows.Scan(&u.Username)
		if err != nil {
			continue
		}

		users = append(users, u)
	}

	return users, nil
}