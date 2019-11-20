package handler

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HelloWorldHander() (method, path string, handler func(c *gin.Context)) {
	return "GET", "/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"ping": "I am alive",
		})
	}

}
