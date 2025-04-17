package handlers

import (
	"auth0_example/platform/authenticator"
	"html/template"
	"net/http"
)

func Login(auth *authenticator.Authenticator) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		t := template.Must(template.ParseFiles("templates/login.html"))
		t.Execute(w, nil)
	}
}
