package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

type MyPage struct {
	Name string
	From string
}

func (my MyPage) SayHello(name string)string{
	return "Hello " + name + ", My name is " + my.Name
}

func TemplateFunction(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.New("FUNCTION").Parse(`{{.SayHello .From }}`))
	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name:"Ega",
		From: "Fahmi",
	})
}

func TestTemplateFunction(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunction(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateFunctionGlobal(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"upper": func(str string)string{
			return strings.ToUpper(str)
		},
	})

	template.Must(t.Parse(`{{upper .From }}`))


	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name:"Ega",
		From: "Fahmi",
	})
}

func TestTemplateFunctionGlobal(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobal(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TemplateFunctionGlobalPipeline(writer http.ResponseWriter, request *http.Request) {
	t := template.New("FUNCTION")
	t.Funcs(map[string]interface{}{
		"upper": func(str string)string{
			return strings.ToUpper(str)
		},
	})

	template.Must(t.Parse(`{{.SayHello .From | upper}}`))


	t.ExecuteTemplate(writer, "FUNCTION", MyPage{
		Name:"Ega",
		From: "Fahmi",
	})
}

func TestTemplateFunctionGlobalPipeline(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateFunctionGlobalPipeline(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}