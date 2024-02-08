package main

import (
	"os"

	"github.com/Edu58/Learn-Go/GinJWTAuth/controllers"
	"github.com/Edu58/Learn-Go/GinJWTAuth/initializers"
	"github.com/gin-gonic/gin"
)

func init() {
	initializers.LoadEnvVariables()
	initializers.ConnectToDB()
	initializers.MigrateModelsToDB()
}

func main() {
	router := gin.Default()

	router.POST("/signup", controllers.SignUp)
	router.POST("/login", controllers.Login)

	port := os.Getenv("PORT")
	router.Run("0.0.0.0:" + port)
}
