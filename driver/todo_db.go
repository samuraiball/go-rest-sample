package driver

import (
	"github.com/jinzhu/gorm"
	"go-rest-sampl/domain"
	"time"
)

type TodoDriver interface {
	GetTodo(todoId string)
}

type TodoDBDriver struct {
	Conn *gorm.DB
}

func (db TodoDBDriver) FindTodoById(todoId string) domain.Todo {
	var todoModel TodoModel

	db.Conn.First(&todoModel, todoId)

	return domain.Todo{
		TodoId:  todoModel.Id,
		Title:   todoModel.Title,
		Content: todoModel.Content,
	}
}

func (db TodoDBDriver) CreateTodo(todo domain.Todo) domain.Todo {

	newTodo := TodoModel{
		Title:   todo.Title,
		Content: todo.Content,
	}
	db.Conn.Create(&newTodo)

	return domain.Todo{
		TodoId:  newTodo.Id,
		Title:   newTodo.Title,
		Content: newTodo.Content,
	}

}

func (db TodoDBDriver) UpdateTodo(todoId int64, todo domain.Todo) domain.Todo {

	var updateTodo TodoModel
	db.Conn.First(&updateTodo, todoId)
	updateTodo.Title = todo.Title
	updateTodo.Content = todo.Content

	db.Conn.Save(updateTodo)

	return domain.Todo{
		TodoId:  updateTodo.Id,
		Title:   updateTodo.Title,
		Content: updateTodo.Content,
	}
}

func (db TodoDBDriver) DeleteTodo(todoId int64) {
	db.Conn.Delete(&TodoModel{Id: todoId})
}

type TodoModel struct {
	Id        int64  `gorm:"primary_key"`
	Title     string `gorm:"title"`
	Content   string `gorm:"content"`
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt *time.Time `sql:"index"`
}
