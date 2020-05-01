package main

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"net/http"
)

type Name struct {
	Name string
}

func mainPageHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "mainPage.html")

}
func mainPageHandler1(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "mainPage.css")
}
func aboutHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "about.html")
}
func mainPagejsHandler(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "mainPage.js")
}
