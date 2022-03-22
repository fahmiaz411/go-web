package goweb

import (
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TemplateAutoEscape(writer http.ResponseWriter, request *http.Request){
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title":"Auto Escape",
		"Body":"<p>Ini adalah body</p><script>alert('anda di heck')</script>",
	})
}

func TestAutoEscape(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	recorder := httptest.NewRecorder()

	TemplateAutoEscape(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}

func TestAutoEscapeServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscape),
	}

	server.ListenAndServe()
}

// not safe

func TemplateAutoEscapeDisable(writer http.ResponseWriter, request *http.Request){
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title":"Auto Escape",
		"Body": template.HTML ("<p>Ini adalah body</p>"),
	})
}

func TestAutoEscapeDisableServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateAutoEscapeDisable),
	}

	server.ListenAndServe()
}
func TemplateXSS(writer http.ResponseWriter, request *http.Request){
	myTemplates.ExecuteTemplate(writer, "post.gohtml", map[string]interface{}{
		"Title":"Auto Escape",
		"Body": template.HTML (request.URL.Query().Get("body")),
	})
}

func TestXSSServer(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(TemplateXSS),
	}

	server.ListenAndServe()
}