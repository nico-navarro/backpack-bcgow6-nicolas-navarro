package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Age    int
	Height int
	Active bool
	Date   string
}

var users []User = []User{}

func Create(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(404, gin.H{
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
