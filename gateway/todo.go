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

func (d TodoGateway) CreateTodo(todo domain.Todo) domain.Todo {
	return d.TodoDriver.CreateTodo(todo)
}

func (d TodoGateway) UpdateTodo(todoId int64, todo domain.Todo) domain.Todo {
	return d.TodoDriver.UpdateTodo(todoId, todo)
}

func (d TodoGateway) DeleteTodo(todoId int64) {
	d.TodoDriver.DeleteTodo(todoId)
}
