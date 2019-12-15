package midleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func LogMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("User-Agent")
		log.Printf("info: there is an accsecc : " + header)
	}
}
