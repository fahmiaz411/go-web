package goweb

import (
	"bytes"
	_ "embed"
	"fmt"
	"io"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"
)

func UploadForm(writer http.ResponseWriter, request *http.Request){
	myTemplates.ExecuteTemplate(writer, "upload_form.gohtml", nil)
}

func Upload(writer http.ResponseWriter, request *http.Request){
	request.ParseMultipartForm(60 << 20)
	file, fileHeader, _ := request.FormFile("file")
	fileDestination, _ := os.Create("./resources/" + fileHeader.Filename)
	_, err := io.Copy(fileDestination, file)
	
	if err != nil {
		panic(err)
	}

	name := request.PostFormValue("name")
	myTemplates.ExecuteTemplate(writer, "upload.success.gohtml", map[string]interface{}{
		"Name": name,
		"File": "/static/" + fileHeader.Filename,
	})
}

func TestUploadForm(t *testing.T) {
	mux:= http.NewServeMux()
	mux.HandleFunc("/", UploadForm)
	mux.HandleFunc("/upload", Upload)
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("./resources"))) )
	
	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	server.ListenAndServe()
}

//go:embed resources/KT.png
var uploadFileTest []byte
func TestUploadFile(t *testing.T) {
	bodyPart := new(bytes.Buffer)

	writer := multipart.NewWriter(bodyPart)
	writer.WriteField("name", "Fahmi Aziz")

	file, _ := writer.CreateFormFile("file", "CONTOH.png")
	file.Write(uploadFileTest)
	writer.Close()


	request := httptest.NewRequest("POST", "http://localhost:8080/", bodyPart)
	request.Header.Set("Content-type",writer.FormDataContentType())

	recorder := httptest.NewRecorder()

	Upload(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)
	fmt.Println(string(body))
}