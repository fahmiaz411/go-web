package goweb

import (
	"embed"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

//go:embed templates/*.gohtml
var allTemplate embed.FS

var myTemplates = template.Must(template.ParseFS(allTemplate, "templates/*.gohtml"))

func TemplateCaching(writer http.ResponseWriter, request *http.Request){
	myTemplates.ExecuteTemplate(writer, "simple.gohtml", "Hello Caching")

}

func TestTemplateCaching(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateCaching(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}