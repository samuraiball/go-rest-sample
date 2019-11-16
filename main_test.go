package main

import (
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
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

func TestGetリクエストでTodoListを取得できる(t *testing.T) {

	router := GinMainEngine()
	w := performRequest(router, "GET", "/todos")

	expected := `
{
   "toodo-lists":[
      {
         "title":"hoge1",
         "content":"piyo1"
      },
      {
         "title":"hoge2",
         "content":"piyo2"
      }
   ]
}`
	assert.Equal(t, http.StatusOK, w.Code)
	assert.JSONEq(t, expected, w.Body.String())
}
