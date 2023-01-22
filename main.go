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
		id := c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
			"accessTime": time,
			"reqestedId": id,
		})
	})

	engine.POST("/", func(c *gin.Context) {
		time := time.Now()
		name := c.PostForm("name")
		age := c.PostForm("age")
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"accessTime": time,
			"name":       name,
			"age":        age,
		})
	})

	engine.GET("/user", func(c *gin.Context) {
		id := c.Query("id")
		c.JSON(http.StatusOK, gin.H{
			"message":    "hello world",
			"User-Agent": ua,
			"id":         id,
		})
	})
	engine.Run(":3000")
}
