package todolist

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func getTodos(c *gin.Context) {
	todos := []todo{{"hoge1", "piyo1"}, {"hoge2", "piyo2"}}
	c.JSON(http.StatusOK, gin.H{
		"todo-list": todos,
	})
}

type todo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}
