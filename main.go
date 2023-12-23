package main

import (
	"database/sql"
	"fmt"
	"net/http"

	_ "github.com/mattn/go-sqlite3"
)

var db, err = sql.Open("sqlite3", "db.sqlite3")

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
	mux.HandleFunc("/api/auth/register", register)
	mux.HandleFunc("/api/auth/exist", userexist)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	fmt.Println("Server started at http://" + addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err)
	}
}
