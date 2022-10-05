package main

import (
	"go-web/cmd/server/handler"
	"go-web/internal/users"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	userRepository := users.NewRepository()
	userService := users.NewService(userRepository)
	userController := handler.NewUserController(userService)

	router := gin.Default()

	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetAll)
	userRouter.POST("/", userController.Store)
	userRouter.PUT("/:id", userController.Update)
	userRouter.DELETE("/:id", userController.Delete)
	userRouter.PATCH("/:id", userController.Patch)

	router.Run()
}
