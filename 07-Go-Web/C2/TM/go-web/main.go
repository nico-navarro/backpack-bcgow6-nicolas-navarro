package main

import (
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type User struct {
	Id     int
	Name   string `json:"name" binding:"required"`
	Email  string `binding:"required"`
	Age    int    `binding:"required"`
	Height int    `binding:"required"`
	Active bool   `binding:"required"`
	Date   string `binding:"required"`
}

var users []User = []User{}

func Create(c *gin.Context) {
	var user User

	token := c.GetHeader("token")
	if token != "123456" {
		c.JSON(401, gin.H{
			"error": "no tiene permisos para realizar la petición solicitada",
		})
		return
	}

	if err := c.ShouldBindJSON(&user); err != nil {
		var ve validator.ValidationErrors
		if errors.As(err, &ve) { //si es de validation errors entra aqui, lo personalizamos
			errorMessages := []string{}
			for _, e := range err.(validator.ValidationErrors) {
				errorMessage := fmt.Sprintf("el campo %s es %s", e.Field(), e.ActualTag())
				errorMessages = append(errorMessages, errorMessage)
			}
			c.JSON(http.StatusBadRequest, errorMessages)
			return
		}
		c.JSON(http.StatusBadRequest, err.Error()) //el resto de errors queda default
		return
	}

	if len(users) > 0 {
		user.Id = users[len(users)-1].Id + 1
	} else {
		user.Id = 1
	}
	users = append(users, user)

	c.JSON(200, user)
}

func main() {
	// ROUTER
	router := gin.Default()

	// example
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Nico!",
		})
	})

	// Create User
	router.POST("/users", Create)

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
