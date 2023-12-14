package main

import (
	"fmt"
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
	tmpl, _ = template.ParseFiles()
}

func main() {
	addr := "localhost:80"
	mux := http.NewServeMux()
	mux.HandleFunc("/", index)
	mux.HandleFunc("/favicon.ico", favicon)
	fileServer := http.FileServer(http.Dir("./static/"))
	mux.Handle("/static/", http.StripPrefix("/static", fileServer))
	fmt.Println("Server started at " + addr)
	err := http.ListenAndServe(addr, mux)
	if err != nil {
		fmt.Println(err)
	}
}
