package http

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/backeseduardo/rinha-backend-2023-golang/internal/person"
	"github.com/gin-gonic/gin"
)

type HttpServerOpts struct {
	PersonRepository person.Repository
}

var opts *HttpServerOpts

func NewServer(o *HttpServerOpts) {
	opts = o

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
	var p person.Person

	err := json.NewDecoder(c.Request.Body).Decode(&p)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	id, err := opts.PersonRepository.Insert(p)
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
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	p, err := opts.PersonRepository.FindById(id)
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
	t := c.Query("t")

	p, err := opts.PersonRepository.FindByTerm(t)
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
	count, err := opts.PersonRepository.Count()
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, gin.H{
			"status":  500,
			"message": err.Error(),
		})
		return
	}

	c.String(http.StatusOK, strconv.Itoa(count))
}
