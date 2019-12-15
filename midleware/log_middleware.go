package midleware

import (
	"github.com/gin-gonic/gin"
	"log"
)

func FooMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("info: fooBefore")

		c.Next()

		log.Printf("info: fooAfter")
	}
}

func BarMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		log.Printf("info: barBefore")

		c.Next()

		log.Printf("info: barAfter")
	}
}
