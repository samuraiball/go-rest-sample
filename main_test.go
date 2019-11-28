package main

import (
	"github.com/stretchr/testify/assert"
	"go-rest-sampl/db"
	"go-rest-sampl/driver"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func performRequest(r http.Handler, method, path string) *httptest.ResponseRecorder {
	req, _ := http.NewRequest(method, "http://localhost:8080/api"+path, nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, req)
	return w
}

func TestPingのPathへアクセスすると文字列を返す(t *testing.T) {

	router := GinMainEngine()
	w := performRequest(router, "GET", "/ping")

	assert.Equal(t, http.StatusOK, w.Code)

	expected := `{"ping": "I am alive"}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}

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

func TestGetリクエストでTodoListを取得できる(t *testing.T) {

	todoId := "2"
	router := GinMainEngine()
	w := performRequest(router, "GET", "/todo/"+todoId)

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
