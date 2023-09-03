package main

import (
	"log"
	"umi/router"

	"github.com/gin-gonic/gin"
)

func main() {
	log.Println("ServerStarting...")
	engine := gin.Default()
	engine.Use(func(c *gin.Context) {
		log.Println("All Request Middleware")
		c.Next()
	})
	router.SetRoutes(engine)
	engine.Run(":3000")
}
