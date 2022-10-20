package main

import (
	"C2/cmd/server/handler"
	"C2/internal/users"
	"C2/pkg/store"
	"fmt"
	"os"

	"C2/docs"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func DummyMiddleware(c *gin.Context) {
	fmt.Println("Im a dummy!")
	// Pass on to the next-in-chain
	c.Next()
}

func AnotherMiddleware(c *gin.Context) {
	fmt.Println("Im another dummy!")
	// Pass on to the next-in-chain
	c.Next()
}

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

	router.Use(DummyMiddleware) // se ejecuta en todas las que siguen hacia abajo

	userRouter := router.Group("/users")
	userRouter.GET("/", userController.GetAll)
	userRouter.POST("/", AnotherMiddleware, userController.Store)    //AnotherMiddleware se ejecuta antes de Store
	userRouter.PUT("/:id", userController.Update, AnotherMiddleware) //AnotherMiddleware se ejecuta después de Update
	userRouter.DELETE("/:id", userController.Delete)
	userRouter.PATCH("/:id", userController.Patch)

	router.Run()
}
