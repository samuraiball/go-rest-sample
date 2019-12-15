package main

import (
	"github.com/comail/colog"
	"github.com/gin-gonic/gin"
	"go-rest-sampl/db"
	"go-rest-sampl/driver"
	"go-rest-sampl/handler"
	"go-rest-sampl/midleware"
	"log"
)

func GinMainEngine() *gin.Engine {
	r := gin.Default()

	apiGroup := r.Group("/api")
	apiGroup.Use(midleware.LogMiddleware())
	apiGroup.Handle(handler.HelthCheckHandler())
	apiGroup.Handle(handler.GetTodoHandler())
	apiGroup.Handle(handler.PostTodoHandler())
	apiGroup.Handle(handler.PutTodoHandler())
	apiGroup.Handle(handler.DeleteTodoHandler())

	return r
}

func main() {

	logSetting()

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

func logSetting() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
}
