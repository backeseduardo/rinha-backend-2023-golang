package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/backeseduardo/rinha-backend-2023-golang"
	"github.com/gin-gonic/gin"
)

func NewServer() {
	r := gin.Default()
	registerRoutes(r)
	r.Run()
}

func registerRoutes(r *gin.Engine) {
	r.POST("/pessoas", HandlePostPessoas)
	r.GET("/pessoas/:id", HandleGetPessoaById)
	r.GET("/pessoas", HandleGetPessoas)
	r.GET("/contagem-pessoas", HandleGetContagemPessoas)
	r.GET("/:id", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": c.Param("id"),
		})
	})
}
func HandlePostPessoas(c *gin.Context) {
	var pessoaBody rinha.Pessoa

	err := json.NewDecoder(c.Request.Body).Decode(&pessoaBody)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	id, err := rinha.InsertPerson(&pessoaBody)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.Header("Location", fmt.Sprintf("/pessoas/%d", id))
	c.JSON(http.StatusCreated, gin.H{"id": id})
}

func HandleGetPessoaById(c *gin.Context) {
	id := c.Param("id")

	p, err := rinha.GetPerson(id)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func HandleGetPessoas(c *gin.Context) {
	term := c.Query("t")

	p, err := rinha.GetPersonsByTerm(term)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, p)
}

func HandleGetContagemPessoas(c *gin.Context) {
	count, err := rinha.CountPersons()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, strconv.Itoa(count))
}
