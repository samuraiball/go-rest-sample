package port

import "go-rest-sampl/domain"

type TodoPort interface {
	FindTodoById(todoId string) domain.Todo
	CreateTodo(todo domain.Todo)
}
