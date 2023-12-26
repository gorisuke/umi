package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the umi API!!",
		})
	})

	engine.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}

func SetAppRoutes(application *gin.RouterGroup) {
	application.GET("/b", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "inside"})
	})
}
