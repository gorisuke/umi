package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthUmi(c *gin.Context) bool {
	log.Println("Use Middleware!!!!!!! Another Files")
	return true
}
