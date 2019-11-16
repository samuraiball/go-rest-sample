package helloworld

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func WorldHandler(router *gin.RouterGroup) {
	router.GET("/ping", HelloWorld)
}

func HelloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"ping": "I am alive",
	})
}
