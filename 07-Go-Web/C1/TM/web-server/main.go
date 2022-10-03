package main

import "github.com/gin-gonic/gin"

func main() {
	// ROUTER
	router := gin.Default()

	// HANDLER
	router.GET("/hello-world", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "Hello World!",
		})
	})
	// Corremos nuestro servidor sobre el puerto 8080
	router.Run()
}
