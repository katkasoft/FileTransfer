package main

import (
	"fmt"
	"net/http"
)

func register(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	password := r.URL.Query().Get("password")
	_, err := db.Exec("insert into users values ('" + email + "', '" + password + "')")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Fprintln(w, "success")
}
