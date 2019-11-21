package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-sampl/handler"
)

func GinMainEngine() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	apiGroup.Handle(handler.TodosHandler())
	apiGroup.Handle(handler.HelloWorldHander())

	return r
}

func main() {

	router := GinMainEngine()
	if err := router.Run(); err != nil {
		panic("アプリケーションの起動に失敗しました。")
	}
}
