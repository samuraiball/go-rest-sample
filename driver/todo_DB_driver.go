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
	db.Conn.AutoMigrate(&domain.Todo{})
	var todo domain.Todo
	db.Conn.First(todo, 1)
	return todo
}
