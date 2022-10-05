package main

import (
	"go-web/cmd/server/handler"
	"go-web/internal/users"
	"go-web/pkg/store"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	_ = godotenv.Load()

	db := store.NewStore("./products.json")
	userRepository := users.NewRepository(db)
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
