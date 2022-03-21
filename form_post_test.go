package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func FormPost(writer http.ResponseWriter, request *http.Request){
	err := request.ParseForm()
	if err != nil {
		panic(err)
	}

	first_name := request.PostForm.Get("first_name")
	last_name := request.PostForm.Get("last_name")

	fmt.Fprintf(writer, "Hello %s, %s", first_name, last_name)
}

func TestFormPost(t *testing.T) {
	requestBody := strings.NewReader("first_name=fahmi&last_name=aziz")
	request := httptest.NewRequest("POST", "http://localhost:8080/", requestBody)
	request.Header.Add("Content-type", "application/x-www-form-urlencoded")

	recorder := httptest.NewRecorder()

	FormPost(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	fmt.Println(string(body))

}