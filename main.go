package main

import (
	"log"
	"net/http"
	"time"

	"github.com/fatih/color"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.Default()
	ua := ""
	color.Green("Prints text in cyan.")

	engine.Use(func(c *gin.Context) {
		ua = c.GetHeader("User-Agent")
		log.Println("All Request Middleware")
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

	engine.GET("/schedule", protectedMiddleware(), func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "hello world",
		})
	})
	engine.Run(":3000")
}
func protectedMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("before logic")
		c.Next()
		log.Println("after logic")
	}
}
