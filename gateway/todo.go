package gateway

import (
	"go-rest-sampl/domain"
	"go-rest-sampl/driver"
)

type TodoGateway struct {
	TodoGateway driver.TodoDBDriver
}

func (d TodoGateway) GetTodo(todoId string) domain.Todo {
	return d.GetTodo(todoId)
}
