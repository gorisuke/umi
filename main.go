package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""
	// ミドルウェアを使用
	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		fmt.Println("Use Middle Ware")
		c.Next()
	})
	engine.GET("/", func(c *gin.Context) {
		time := time.Now()
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
			"accessTime": time,
		})
	})
	engine.Run(":3000")
}
