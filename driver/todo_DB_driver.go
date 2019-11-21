package driver

import (
	"github.com/jinzhu/gorm"
	"go-rest-sampl/domain"
)

type TodoDriver interface {
	GetTodo(todoId string)
}

type TodoDBDriver struct {
	Conn *gorm.DB
}

func (db TodoDBDriver) GetTodo(todoId string) domain.Todo {
	return domain.Todo{
		Title:   "title",
		Content: "content",
	}
}
