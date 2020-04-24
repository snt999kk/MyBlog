package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"os"
	"time"
)

func main() {
	Mux := mux.NewRouter()
	Mux.HandleFunc("/", mainPageHandler)
	Mux.HandleFunc("/mainPage.css", mainPageHandler1)
	Mux.HandleFunc("/name", getNameHandler)
	Mux.HandleFunc("/about", aboutHandler)
	srv := &http.Server{
		Addr:         ":" + os.Getenv("PORT"),
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      Mux,
	}
	srv.ListenAndServe()
}
