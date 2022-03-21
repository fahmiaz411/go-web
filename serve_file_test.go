package goweb

import (
	_ "embed"
	"fmt"
	"net/http"
	"testing"
)

func ServeFile(writer http.ResponseWriter, request *http.Request){
	if request.URL.Query().Get("name") != ""{
		http.ServeFile(writer, request, "./resources/index.html")
	} else {
		http.ServeFile(writer, request, "./resources/notfound.html")
	}
}

func TestServeFileListener(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFile),
	}

	server.ListenAndServe()
}

//go:embed resources/index.html
var resourceOk string

//go:embed resources/notfound.html
var resourceNotFound string

func ServeFileEmbed(writer http.ResponseWriter, request *http.Request){
	if request.URL.Query().Get("name") != ""{
		fmt.Fprint(writer, resourceOk)
	} else {
		fmt.Fprint(writer, resourceNotFound)
	}
}

func TestServeFileEmbedListener(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(ServeFileEmbed),
	}

	server.ListenAndServe()
}
