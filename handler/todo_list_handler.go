package handler

import (
	"github.com/gin-gonic/gin"
	"go-rest-sampl/db"
	"go-rest-sampl/driver"
	"go-rest-sampl/gateway"
	"go-rest-sampl/usecase"
	"net/http"
)

func TodosHandler() (method, path string, handler func(c *gin.Context)) {
	return "GET", "/todos", func(c *gin.Context) {

		todo := usecase.TodoPort{
			TodoPort: gateway.TodoGateway{
				TodoGateway: driver.TodoDBDriver{
					Conn: db.DB(),
				},
			},
		}

		c.JSON(http.StatusOK, gin.H{
			"todo-list": todo,
		})
	}
}

type todo struct {
	Title   string `json:"title"`
	Content string `json:"content"`
}