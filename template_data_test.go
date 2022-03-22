package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateDataMap(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", map[string]interface{}{
		"Name":  "fahmi",
		"Title": "Sample Test",
		"Addr": map[string]string{
			"Name": "Jl",
		},
	})
}

func TestTemplateDataMap(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataMap(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Address struct {
	Name string
}

type TemplateStruct struct {
	Name  string
	Title string
	Addr  Address
}

func TemplateDataStruct(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/name.gohtml"))
	t.ExecuteTemplate(writer, "name.gohtml", TemplateStruct{
		Name:  "fahmi",
		Title: "Sample Test",
		Addr: Address{
			Name: "Jalan Kp",
		},
	})
}

func TestTemplateDataStruct(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateDataStruct(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}