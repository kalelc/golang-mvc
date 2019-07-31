package main

import (
	"fmt"
	"net/http"

	"github.com/kalelc/golang-mvc/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func main() {
	homeView = views.NewView("views/home.html")
	contactView = views.NewView("views/contact.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	fmt.Println("localhost:3000")
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := homeView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	if err := contactView.Template.Execute(w, nil); err != nil {
		panic(err)
	}
}
