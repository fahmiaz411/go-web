package goweb

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httptest"
	"testing"
)

func SetCookie(writer http.ResponseWriter, request *http.Request){
	fazcookie := new(http.Cookie)
	fazcookie.Name = "X-FAZ-Name"
	fazcookie.Value = request.URL.Query().Get("name")
	fazcookie.Path = "/"

	http.SetCookie(writer, fazcookie)
	fmt.Fprint(writer, "Success create cookie")
}

func GetCookie(writer http.ResponseWriter, request *http.Request){
	cookie, err := request.Cookie("X-FAZ-Name")

	if err != nil {
		fmt.Fprint(writer, "No cookie")
	} else {
		name := cookie.Value
		fmt.Fprintf(writer, "Cookie is %s", name)
	}
}

func TestCookie(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/set-cookie", SetCookie)
	mux.HandleFunc("/get-cookie", GetCookie)

	server := http.Server{
		Addr: "localhost:8080",
		Handler: mux,
	}

	err := server.ListenAndServe()
	if err != nil {
		panic(err)
	}
}

func TestSetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/?name=faz", nil)
	recorder := httptest.NewRecorder()

	SetCookie(recorder, request)

	cookies := recorder.Result().Cookies()

	for _, cookie := range cookies{
		fmt.Printf("Cookie %s:%s\n", cookie.Name, cookie.Value)
	}
}
func TestGetCookie(t *testing.T) {
	request := httptest.NewRequest("GET", "http://localhost:8080/", nil)
	cookie := new(http.Cookie)
	cookie.Name = "X-FAZ-Name"
	cookie.Value = "faz"
	request.AddCookie(cookie)

	recorder := httptest.NewRecorder()

	GetCookie(recorder, request)

	body, _ := io.ReadAll(recorder.Result().Body)

	fmt.Println(string(body))

}