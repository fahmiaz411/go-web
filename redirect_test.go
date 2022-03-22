package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

func RedirectTo(writer http.ResponseWriter, request *http.Request){
	fmt.Fprint(writer, "Hello redirect")
}

func RedirectFrom(writer http.ResponseWriter, request *http.Request){
	// logic
	// http.Redirect(writer, request, "https://www.facebook.com", http.StatusTemporaryRedirect)
	http.Redirect(writer, request, "/redirect-to", http.StatusTemporaryRedirect)
}

func TestRedirect(t *testing.T){
	mux := http.NewServeMux()
	mux.HandleFunc("/redirect-from", RedirectFrom)
	mux.HandleFunc("/redirect-to", RedirectTo)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}