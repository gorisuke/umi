package auth

import (
	"log"

	"github.com/gin-gonic/gin"
)

func AuthUmi(c *gin.Context) bool {
	log.Printf("test")
	return true
}

func JwtAuthentication() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Println("JWT Auth")
		flag := false
		if flag {
			c.Next()
		} else {
			c.JSON(401, gin.H{"message": "Invalid Access"})
			c.Abort()
		}
	}
}
