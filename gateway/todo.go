package gateway

import (
	"go-rest-sampl/domain"
	"go-rest-sampl/driver"
)

type TodoGateway struct {
	TodoDriver driver.TodoDBDriver
}

func (d TodoGateway) GetTodo(todoId string) domain.Todo {
	return d.TodoDriver.GetTodo(todoId)
}
