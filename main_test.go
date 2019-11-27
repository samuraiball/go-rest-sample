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

	code := m.Run()
	os.Exit(code)
}

func TestGetリクエストでTodoListを取得できる(t *testing.T) {

	//todoId := "1"
	router := GinMainEngine()
	w := performRequest(router, "GET", "/todos")

	expected := `
{
   "todo-list":
      {
         "todo_id": 1,
         "title":"title1",
         "content":"content1"
      }
}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}
