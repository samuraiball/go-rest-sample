package handler

import (
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

func HelthCheckHandler() (method, path string, handler func(c *gin.Context)) {
	return "GET", "/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "I am alive",
		})
		log.Printf("info: ping is called")
	}

}
