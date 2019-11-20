package port

import "go-rest-sampl/domain"

type TodoPort interface {
	GetTodo(todoId string) domain.Todo
	GetTodoList(asset, limit string) []domain.Todo
	CreateTodo(todo domain.Todo) domain.Todo
}
