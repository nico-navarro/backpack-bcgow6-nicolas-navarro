package main

import (
	"encoding/json"
	"io/ioutil"

	"github.com/gin-gonic/gin"
)

type User struct {
	Id     int
	Name   string
	Email  string
	Age    int
	Height string
	Active bool
	Date   string
}

type Users struct {
	Users []User
}

func GetAll(c *gin.Context) {
	file, _ := ioutil.ReadFile("users.json")

	users := []User{}

	_ = json.Unmarshal([]byte(file), &users)

	c.JSON(200, users)
}

func main() {
	// ROUTER
	router := gin.Default()

	// HANDLER
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Nico!",
		})
	})

	// Get All Users
	router.GET("/users", GetAll)

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
