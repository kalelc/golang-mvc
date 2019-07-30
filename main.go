package main

import (
	"golang-mvc/views"
	"html/template"
	"net/http"

	"github.com/gorilla/mux"
)

var homeView *views.View

func main() {
	var err error
	homeTemplate, err = template.ParseFiles(
		"views/home.gohtml",
		"views/layouts/top.gohtml")
	if err != nil {
		panic(err)
	}

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeTemplate.Execute(w, nil); err != nil {
		panic(err)
	}
}
