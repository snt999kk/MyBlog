package main

import (
	//"encoding/json"
	//"fmt"
	//"io/ioutil"
	"net/http"
	"html/template"
	)
type Name struct{
	Name string
}
func mainPageHandler(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "mainPage.html")

}
func mainPageHandler1(w http.ResponseWriter, r *http.Request){
	http.ServeFile(w, r, "mainPage.css")
}
func getNameHandler(w http.ResponseWriter, r *http.Request){
	name := Name{r.FormValue("Name")}
	tmpl := template.Must(template.ParseFiles("hello.html"))
	//fmt.Println(name)
	tmpl.Execute(w, name)
}
func aboutHandler(w http.ResponseWriter, r* http.Request){
	http.ServeFile(w, r, "about.html")
}