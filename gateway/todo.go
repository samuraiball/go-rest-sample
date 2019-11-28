package gateway

import (
	"go-rest-sampl/domain"
	"go-rest-sampl/driver"
)

type TodoGateway struct {
	TodoDriver driver.TodoDBDriver
}

func (d TodoGateway) FindTodoById(todoId string) domain.Todo {
	return d.TodoDriver.FindTodoById(todoId)
}

func (d TodoGateway) CreateTodo(todo domain.Todo) {
	d.TodoDriver.CreateTodo(todo)
}
