package main

import (
	"github.com/gin-gonic/gin"
	"go-rest-sampl/db"
	"go-rest-sampl/driver"
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
	db.DB().LogMode(true)

	initDb()

	defer db.CloseDB()

	if err := router.Run(); err != nil {
		panic("アプリケーションの起動に失敗しました。")
	}
}

func initDb() {
	db.DB().DropTableIfExists()
	db.DB().AutoMigrate(&driver.TodoModel{})
}
