package main

import (
	"log"
	"umi/auth"
	"umi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("ServerStarting")
	engine := gin.Default()
	engine.Use(auth.JwtAuthentication())
	router.SetRoutes(engine)
	engine.Run(":3000")
}
