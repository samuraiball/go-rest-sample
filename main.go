package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func GinMainEngine() *gin.Engine {
	router := gin.Default()
	router.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"hello": "world",
		})
	})
	return router
}

func main() {
	router := GinMainEngine()
	if err := router.Run(); err != nil {
		panic("アプリケーションの起動に失敗しました。")
	}
}
