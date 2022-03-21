package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func RequestResponseHeader(writer http.ResponseWriter, request *http.Request) {
	contentType := request.Header.Get("content-type")

	writer.Header().Add("X-Powered-By", "FAZZ")
	
	fmt.Fprint(writer, contentType)
}

func TestHeader (t *testing.T){
	request := httptest.NewRequest("POST","http://localhost:8080/", nil)
	request.Header.Add("Content-type", "application/json")
	
	recorder := httptest.NewRecorder()

	RequestResponseHeader(recorder, request)

	response := recorder.Result()
	body, _ := io.ReadAll(response.Body)

	poweredBy := response.Header.Get("X-Powered-By")

	fmt.Println(string(body), poweredBy)
}