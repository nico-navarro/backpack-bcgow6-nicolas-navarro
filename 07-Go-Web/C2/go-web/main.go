package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id     int
	Name   string `binding:"required"`
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
		c.JSON(400, gin.H{
			"error": err.Error(),
		})
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
