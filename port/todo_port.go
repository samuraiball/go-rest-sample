package port

import "go-rest-sampl/domain"

type TodoPort interface {
	FindTodoById(todoId string) domain.Todo
	CreateTodo(todo domain.Todo) domain.Todo
	UpdateTodo(todoId int64, todo domain.Todo) domain.Todo
	DeleteTodo(todoId int64)
}
