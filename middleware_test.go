package goweb

import (
	"fmt"
	"net/http"
	"testing"
)

type LogMiddleware struct {
	Handler http.Handler
}

func (middleware *LogMiddleware)ServeHTTP(writer http.ResponseWriter, request *http.Request){
	// middleware here
	fmt.Println("middle")
	middleware.Handler.ServeHTTP(writer, request)
}

type ErrorHandler struct {
	Handler http.Handler
}

func handleError (w http.ResponseWriter){
	err := recover()
		if err != nil {
			fmt.Println("Error")
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "Error: %s", err)
		}
}

func(handler *ErrorHandler)ServeHTTP(w http.ResponseWriter, r *http.Request){
	defer handleError(w)
	handler.Handler.ServeHTTP(w, r)
}

func TestMiddleware(t *testing.T) {
	mux := http.NewServeMux()
	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request){
		// fmt.Println("handler execute")
		fmt.Fprint(w, "Hello middleware")
	})
	mux.HandleFunc("/foo", func(w http.ResponseWriter, r *http.Request){
		// fmt.Println("handler execute")
		fmt.Fprint(w, "Hello Foo")
	})
	mux.HandleFunc("/panic", func(w http.ResponseWriter, r *http.Request){
		// fmt.Println("handler execute")
		// fmt.Fprint(w, "Hello Foo")
		panic("ups")

	})

	logMiddle := &LogMiddleware{
		Handler: mux,
	}

	errorHandler := &ErrorHandler{
		Handler: logMiddle,
	}

	// logMiddle := new(LogMiddleware)
	// logMiddle.Handler = mux

	server := http.Server{
		Addr: "localhost:8080",
		Handler: errorHandler,
	}

	server.ListenAndServe()

}