package driver

import "go-rest-sampl/domain"

func FindTodo(todoId string) domain.Todo {
	return domain.Todo{
		Title:   "title",
		Content: "content",
	}
}

func GetTodo(todoId string) {

}
func GetTodoList(asset, limit string) {

}
func CreateTodo(todo domain.Todo) {

}
