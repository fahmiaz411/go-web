package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

// type Address struct {
// 	Name string
// }

// type TemplateStruct struct {
// 	Name  string
// 	Title string
// 	Addr  Address
// }

func TemplateActionIf(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/if.gohtml"))
	t.ExecuteTemplate(writer, "if.gohtml", TemplateStruct{
		Name:  "fahmi",
		Addr: Address{
			Name: "Jalan Kp",
		},
	})
}

func TestTemplateActionIf(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionIf(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Comparator struct {
	Value int
	T *TemplateStruct
}

func TemplateActionComparator(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/comparator.gohtml"))
	t.ExecuteTemplate(writer, "comparator.gohtml", Comparator{
		Value: 10,
		T: &TemplateStruct{
			Name:  "fahmi",
			Addr: Address{
				Name: "Jalan Kp",
			},
		},
	})
}

func TestTemplateActionComparator(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionComparator(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

type Hobby struct {
	Hobbies []string
	T *TemplateStruct
}

func TemplateActionRange(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/range.gohtml"))
	t.ExecuteTemplate(writer, "range.gohtml", Hobby{
		Hobbies: []string{
			"Memasak",
			"Berenang",
		},
		T: &TemplateStruct{
			Name:  "fahmi",
			Addr: Address{
				Name: "Jalan Kp",
			},
		},
	})
}

func TestTemplateActionRange(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionRange(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}
func TemplateActionWith(writer http.ResponseWriter, request *http.Request) {
	t := template.Must(template.ParseFiles("./templates/with.gohtml"))
	t.ExecuteTemplate(writer, "with.gohtml", Hobby{
		T: &TemplateStruct{
			Name:  "fahmi",
			Addr: Address{
				Name: "Jalan Kp",
			},
		},
	})
}

func TestTemplateActionWith(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateActionWith(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}