package usecase

import (
	"go-rest-sampl/domain"
	"go-rest-sampl/port"
)

type useCase struct {
	port port.TodoPort
}

func GetTodo(todoId string) domain.Todo {
	return domain.Todo{
		Title:   "",
		Content: "",
	}
}

func GetTodoList(asset, limit string) []domain.Todo {
	return nil
}

func CreateTodo(todo domain.Todo) domain.Todo {
	return domain.Todo{
		Title:   "",
		Content: "",
	}
}
