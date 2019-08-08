package controllers

import (
	"fmt"
	"net/http"

	"github.com/kalelc/golang-mvc/views"
)

func NewUsers() *Users {
	return &Users{
		NewView: views.NewView("bootstrap", "views/users/new.html"),
	}
}

func (u *Users) New(w http.ResponseWriter, r *http.Request) {
	if err := u.NewView.Render(w, nil); err != nil {
		panic(err)
	}
}

func (u *Users) Create(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		panic(err)
	}
	fmt.Println(w, r.PostForm["email"])
	fmt.Println(w, r.PostForm["password"])
}

type Users struct {
	NewView *views.View
}
