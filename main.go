package main

import (
	"fmt"
	"net/http"

	"github.com/kalelc/golang-templates/views"

	"github.com/gorilla/mux"
)

var homeView *views.View
var contactView *views.View

func main() {
	homeView = views.NewView("bootstrap", "views/home.html")
	contactView = views.NewView("bootstrap", "views/contact.html")

	r := mux.NewRouter()
	r.HandleFunc("/", home)
	r.HandleFunc("/contact", contact)

	fmt.Println("localhost:3000")
	http.ListenAndServe(":3000", r)
}

func home(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")
	must(homeView.Render(w, nil))
}

func contact(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/html")

	must(contactView.Render(w, nil))
}

func must(err error) {
	if err != nil {
		panic(err)
	}
}
