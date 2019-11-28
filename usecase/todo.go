package usecase

import (
	"go-rest-sampl/domain"
	"go-rest-sampl/port"
)

type TodoPort struct {
	TodoPort port.TodoPort
}

func (port TodoPort) FindTodoById(todoId string) domain.Todo {
	return port.TodoPort.FindTodoById(todoId)
}

func (port TodoPort) CreateTodo(todo domain.Todo) domain.Todo {
	return port.TodoPort.CreateTodo(todo)
}
