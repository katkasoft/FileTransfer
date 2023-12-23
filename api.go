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
		fmt.Fprintln(w, "error")
		return
	}
	cookie := &http.Cookie{
		Name:  "email",
		Value: email,
		Path:  "/",
	}
	http.SetCookie(w, cookie)
	fmt.Fprintln(w, "success")
}

func userexist(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	var count int
	row := db.QueryRow("SELECT COUNT(*) FROM users WHERE email = ?", email)
	err := row.Scan(&count)
	if err != nil {
		fmt.Fprint(w, err)
	}
	if count > 0 {
		fmt.Fprint(w, "true")
	} else {
		fmt.Fprint(w, "false")
	}
}
