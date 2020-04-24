package main

import (
	"github.com/gorilla/mux"
	"net/http"
	"time"
)
func main(){
	Mux := mux.NewRouter()
	Mux.HandleFunc("/", mainPageHandler)
	Mux.HandleFunc("/mainPage.css", mainPageHandler1)
	Mux.HandleFunc("/name", getNameHandler)
	Mux.HandleFunc("/about", aboutHandler)
	srv := &http.Server{
		Addr:":8080",
		ReadTimeout: 5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:Mux,
	}
	srv.ListenAndServe()
}
