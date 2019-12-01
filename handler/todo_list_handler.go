package handler

import (
	"github.com/gin-gonic/gin"
	"go-rest-sampl/db"
	"go-rest-sampl/domain"
	"go-rest-sampl/driver"
	"go-rest-sampl/gateway"
	"go-rest-sampl/usecase"
	"net/http"
	"strconv"
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

		c.JSON(http.StatusOK, gin.H{"todo-list": todo.FindTodoById(c.Param("id"))})
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

		c.JSON(http.StatusCreated, gin.H{"todo-list": todo.CreateTodo(newTodo)})
	}
}

func PutTodoHandler() (method, path string, handler func(c *gin.Context)) {
	return "PUT", "todo/:id", func(c *gin.Context) {

		todo := usecase.TodoPort{
			TodoPort: gateway.TodoGateway{
				TodoDriver: driver.TodoDBDriver{
					Conn: db.DB(),
				},
			},
		}

		// FIXME: error handling
		targetId, _ := strconv.ParseInt(c.Param("id"), 10, 64)

		var updateTodo domain.Todo
		//FIXME: error handling
		if err := c.BindJSON(&updateTodo); err != nil {
			panic(err)
		}

		c.JSON(http.StatusOK, gin.H{"todo-list": todo.UpdateTodo(targetId, updateTodo)})
	}
}
