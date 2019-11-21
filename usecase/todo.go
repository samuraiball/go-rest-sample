package usecase

import (
	"go-rest-sampl/domain"
	"go-rest-sampl/port"
)

type TodoPort struct {
	TodoPort port.TodoPort
}

func (port TodoPort) GetTodo(todoId string) domain.Todo {
	todo := port.TodoPort.GetTodo(todoId)
	return todo
}
