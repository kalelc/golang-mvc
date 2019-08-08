package controllers

import (
	"net/http"

	"github.com/gorilla/schema"
)

func parseForm(r *http.Request, dst interface{}) error {
	if err := r.ParseForm(); err != nil {
		return err
	}
	dec := schema.NewDecoder()
	form := SignupForm{}
	if err := dec.Decode(&form, r.PostForm); err != nil {
		return err
	}
	return nil
}
