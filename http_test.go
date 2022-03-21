package goweb

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"
)

func HelloHandler(writer http.ResponseWriter, request http.Request) {
	fmt.Fprintln(writer, "hello")
}

func TestHttp(t *testing.T){
  request := httptest.NewRequest(http.MethodGet, "http://localhost:8080/hello", nil)
  recorder := httptest.NewRecorder()
}