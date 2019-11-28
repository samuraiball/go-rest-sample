package handler

import (
	"github.com/gin-gonic/gin"
	"go-rest-sampl/db"
	"go-rest-sampl/domain"
	"go-rest-sampl/driver"
	"go-rest-sampl/gateway"
	"go-rest-sampl/usecase"
	"net/http"
)

func GetTodoHandler() (method, path string, handler func(c *gin.Context)) {
	return "GET", "/todo/:id", func(c *gin.Context) {

		todo := usecase.TodoPort{
			TodoPort: gateway.TodoGateway{
				TodoDriver: driver.TodoDBDriver{
					Conn: db.DB(),
				},
			},
		}

		c.JSON(http.StatusOK, gin.H{
			"todo-list": todo.FindTodoById(c.Param("id")),
		})
	}
}

func PostTodoHandler() (method, path string, handler func(c *gin.Context)) {
	return "POST", "todo", func(c *gin.Context) {

		todo := usecase.TodoPort{
			TodoPort: gateway.TodoGateway{
				TodoDriver: driver.TodoDBDriver{
					Conn: db.DB(),
				},
			},
		}

		var newTodo domain.Todo

		//FIXME: error handling
		if err := c.BindJSON(&newTodo); err != nil {
			panic(err)
		}

		todo.CreateTodo(newTodo)

		c.JSON(http.StatusCreated, nil)
	}
}
