package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func TestHandler(t *testing.T){
	var handler http.HandlerFunc = func(writer http.ResponseWriter, request *http.Request){
       fmt.Fprint(writer, "Hello world")
	}

	server := http.Server{
		Addr: "localhost:8080",
		Handler: handler,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestServeMux(t *testing.T){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request){
      fmt.Fprint(writer, "hello world")
	})
	mux.HandleFunc("/hi", func (writer http.ResponseWriter, request *http.Request){
      fmt.Fprint(writer, "hi you")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}
func TestRequest(t *testing.T){
	mux := http.NewServeMux()

	mux.HandleFunc("/", func (writer http.ResponseWriter, request *http.Request){
      fmt.Fprintln(writer, request.Method)
      fmt.Fprintln(writer, request.RequestURI)
	})
	mux.HandleFunc("/hi", func (writer http.ResponseWriter, request *http.Request){
      fmt.Fprint(writer, "hi you")
	})

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}