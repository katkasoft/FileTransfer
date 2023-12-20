package main

import (
	"database/sql"
	"fmt"
	"html/template"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db, err = sql.Open("sqlite3", "db.sqlite3")

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

func register(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	db.Exec("insert into users values ('" + email + "', '" + password + "')")
	fmt.Fprintln()
}

func main() {
	if err != nil {
		fmt.Println(err)
	}
	db.Exec("create table if not exists users (email text, password text)")
	addr := "localhost:80"
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/favicon.ico", favicon)
	mux.HandleFunc("/auth/register", registerview)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	fmt.Println("Server started at http://" + addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err)
	}
}
