package middlewares

import (
	"github.com/gin-gonic/gin"
)

// CorsMiddleware handles cors
func CorsMIddleware() gin.HandlerFunc{
	return func(c *gin.Context){
		c.Writer.Header().Set("Access-Control-Allow-Origin", "*")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "POST, GET, OPTIONS, PUT, DELETE")
		c.Writer.Header().Set("Access-Control-Allow-Origin", "Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")

		if c.Request.Method == "OPTIONS"{
			c.AbortWithStatus(204)
			return
		}

		c.Next()
	}
}