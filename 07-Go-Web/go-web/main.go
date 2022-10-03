package main

import (
	"encoding/json"
	"io/ioutil"
	"strconv"

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

func GetAll(c *gin.Context) {
	file, _ := ioutil.ReadFile("users.json")
	users := []User{}
	_ = json.Unmarshal([]byte(file), &users)

	result := []User{}

	name := c.Query("name")
	email := c.Query("email")
	age, _ := strconv.Atoi(c.Query("age"))

	for _, user := range users {

		conditionName := true
		if name != "" {
			conditionName = user.Name == name
		}

		conditionEmail := true
		if email != "" {
			conditionEmail = user.Email == email
		}

		conditionAge := true
		if age > 0 {
			conditionAge = user.Age == age
		}

		if conditionName && conditionEmail && conditionAge {
			result = append(result, user)
		}
	}
	c.JSON(200, result)
}

func GetOne(c *gin.Context) {
	file, _ := ioutil.ReadFile("users.json")
	users := []User{}
	_ = json.Unmarshal([]byte(file), &users)

	userWanted := User{}
	status := 404
	id, _ := strconv.Atoi(c.Param("id"))
	for _, user := range users {
		if user.Id == id {
			userWanted = user
			status = 200
			break
		}
	}
	c.JSON(status, userWanted)
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

	// Get All Users
	router.GET("/users", GetAll)

	// Get One User
	router.GET("/users/:id", GetOne)

	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
