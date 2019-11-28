package main

import (
	"github.com/stretchr/testify/assert"
	"go-rest-sampl/db"
	"go-rest-sampl/driver"
	"io"
	"net/http"
	"net/http/httptest"
	"os"
	"strings"
	"testing"
)

func TestMain(m *testing.M) {

	db.DB().DropTableIfExists(&driver.TodoModel{})
	db.DB().AutoMigrate(&driver.TodoModel{})

	db.DB().Create(&driver.TodoModel{
		Title:   "title1",
		Content: "content1",
	})
	db.DB().Create(&driver.TodoModel{
		Title:   "title2",
		Content: "content2",
	})

	code := m.Run()
	os.Exit(code)
}

func TestPingのPathへアクセスすると文字列を返す(t *testing.T) {

	router := GinMainEngine()
	w := performRequest(router, "GET", "/ping", nil)

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"ping": "I am alive"}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}

func TestGetリクエストでTodoListを取得できる(t *testing.T) {

	todoId := "2"
	router := GinMainEngine()
	w := performRequest(router, "GET", "/todo/"+todoId, nil)

	expected := `
{
   "todo-list":
      {
         "todo_id": 2,
         "title":"title2",
         "content":"content2"
      }
}`

	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}

func TestPOSTリクエストでTodoを登録できる(t *testing.T) {

	router := GinMainEngine()

	requestBody := `
{
   "title":"posted_todo", 
   "content":"this is a posted todo task"
}
`

	expectedDbInsert := &driver.TodoModel{
		Id:      3,
		Title:   "posted_todo",
		Content: "this is a posted todo task",
	}

	expectedResponse := `
{
   "todo-list":
      {
         "todo_id": 3,
         "title":"posted_todo",
         "content":"this is a posted todo task"
      }
}
`

	w := performRequest(router, "POST", "/todo", strings.NewReader(requestBody))

	actualDbInsert := &driver.TodoModel{}
	db.DB().Find(actualDbInsert).Where("title=?", "posted_todo")

	assert.Equal(t, expectedDbInsert, actualDbInsert)
	assert.Equal(t, http.StatusCreated, w.Code)
	assert.JSONEq(t, expectedResponse, w.Body.String())
}

func performRequest(r http.Handler, method, path string, body io.Reader) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, "http://localhost:8080/api"+path, body)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}
