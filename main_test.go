package main

import (
	"encoding/json"
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

func TestGETリクエストでHelloWorldの文字列が返却される(t *testing.T) {

	router := GinMainEngine()
	w := performRequest(router, "GET", "/ping")

	var res map[string]string
	err := json.Unmarshal([]byte(w.Body.String()), &res)

	assert.Nil(t, err)
	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "I am alive", res["ping"])
}
