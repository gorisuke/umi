package main

import (
	"log"
	"os"
	"umi/auth"
	"umi/initializers"
	"umi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("ServerStarting")
	initializers.LoadEnv()
	initializers.ConnectToDb()
	engine := gin.Default()
	application := engine.Group("/app")
	application.Use(auth.JwtAuthentication())
	router.SetAppRoutes(application)
	router.SetRoutes(engine)
	engine.Run(os.Getenv("PORT"))
}
