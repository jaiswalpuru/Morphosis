package controllers

import (
	"net/http"
	"text/template"
)

func WelcomeHandler(w http.ResponseWriter, r *http.Request) {
	t, _ := template.ParseFiles("./templates/profile.html")
	t.Execute(w, nil)
}
