package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-sampl/handler/helloworld"
	"go-rest-sampl/handler/todolist"
)

func GinMainEngine() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	apiGroup.Handle(todolist.TodosHandler())
	apiGroup.Handle(helloworld.HelloWorldHander())

	return r
}

func main() {
	router := GinMainEngine()
	if err := router.Run(); err != nil {
		panic("アプリケーションの起動に失敗しました。")
	}
}
