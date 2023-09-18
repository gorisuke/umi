package router

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func SetRoutes(engine *gin.Engine) {
	engine.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"message": "Welcome to the umi API!!!",
		})
	})
}
