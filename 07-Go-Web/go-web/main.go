package main

import "github.com/gin-gonic/gin"

func main() {
	// ROUTER
	router := gin.Default()

	// HANDLER
	router.GET("/hello", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hola Nico!",
		})
	})
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
