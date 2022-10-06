package main

import (
	"go-web/cmd/server/handler"
	"go-web/internal/users"
	"go-web/pkg/store"
	"os"

	"go-web/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @title MELI Bootcamp API
// @version 1.0
// @description This API Handle MELI Products.
// @termsOfService https://developers.mercadolibre.com.ar/es_ar/terminos-y-condiciones
// @contact.name API Support
// @contact.url https://developers.mercadolibre.com.ar/support
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
func main() {
	_ = godotenv.Load()

	db := store.NewStore("./products.json")
	userRepository := users.NewRepository(db)
	userService := users.NewService(userRepository)
	userController := handler.NewUserController(userService)

	router := gin.Default()

	docs.SwaggerInfo.Host = os.Getenv("HOST")
	router.GET("/docs/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetAll)
	userRouter.POST("/", userController.Store)
	userRouter.PUT("/:id", userController.Update)
	userRouter.DELETE("/:id", userController.Delete)
	userRouter.PATCH("/:id", userController.Patch)

	router.Run()
}
