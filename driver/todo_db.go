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
	var todoModel TodoModel

	db.Conn.First(&todoModel, todoId)

	return domain.Todo{
		TodoId:  todoModel.Id,
		Title:   todoModel.Title,
		Content: todoModel.Content,
	}
}

type TodoModel struct {
	Id      int64  `gorm:"primary_key"`
	Title   string `gorm:"title"`
	Content string `gorm:"content"`
}
