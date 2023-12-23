package main

import (
	"html/template"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		tmpl, _ := template.ParseFiles("templates/404.html")
		tmpl.Execute(w, nil)
		return
	}
	tmpl, _ := template.ParseFiles("templates/index.html")
	tmpl.Execute(w, nil)
}

func favicon(w http.ResponseWriter, r *http.Request) {
	http.ServeFile(w, r, "static/favicon.ico")
}

func registerview(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/auth/register.html")
	tmpl.Execute(w, nil)
}

func loginview(w http.ResponseWriter, r *http.Request) {
	tmpl, _ := template.ParseFiles("templates/auth/login.html")
	tmpl.Execute(w, nil)
}
