package port

import "go-rest-sampl/domain"

type TodoPort interface {
	GetTodo(todoId string) domain.Todo
}
