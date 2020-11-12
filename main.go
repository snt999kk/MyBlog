package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	_ "os"
	"time"
)

func main() {
	port := "8081" //os.Getenv("PORT")
	Mux := mux.NewRouter()
	Mux.HandleFunc("/", mainPageHandler)
	Mux.HandleFunc("/mainPage.css", mainPageHandler1)
	Mux.HandleFunc("/about", aboutHandler)
	Mux.HandleFunc("/mainPage.js", mainPagejsHandler)
	Mux.HandleFunc("/news", newsHandler)
	srv := &http.Server{
		Addr:         ":" + port,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
		Handler:      Mux,
	}
	fmt.Println(port)
	srv.ListenAndServe()
}
