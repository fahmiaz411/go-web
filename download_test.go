package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

//content disposition
// inline -- render
// attachment -- no render
// attachment; filename="filename.jpg" -- no render and filename

func DownloadFile(writer http.ResponseWriter, request *http.Request){
	file := request.URL.Query().Get("file")
	if file == ""{
		writer.WriteHeader(http.StatusBadRequest)
		fmt.Fprint(writer, "Bad Request")
		return
	}

	writer.Header().Add("Content-Disposition", "attachment; filename=\""+file+"\"")
	http.ServeFile(writer, request, "./resources/" + file)

}

func TestDownloadFile(t *testing.T) {
	server := http.Server{
		Addr: "localhost:8080",
		Handler: http.HandlerFunc(DownloadFile),
	}

	server.ListenAndServe()
}