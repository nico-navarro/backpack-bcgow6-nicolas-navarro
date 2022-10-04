package main

import (
	"go-web/cmd/server/handler"
	"go-web/internal/users"

	"github.com/gin-gonic/gin"
)

func main() {
	userRepository := users.NewRepository()
	userService := users.NewService(userRepository)
	userController := handler.NewUser(userService)

	router := gin.Default()

	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetAll())
	userRouter.POST("/", userController.Store())

	router.Run()
}
