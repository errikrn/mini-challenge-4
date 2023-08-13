package main

import (
	"fmt"
	"net/http"
	"text/template"
)

type User struct {
	Name      string
	Email     string
	Alamat    string
	Pekerjaan string
	Alasan    string
}

var users = map[string]User{
	"fitri@mail.com": User{Name: "Fitri", Email: "fitri@mail.com", Alamat: "Jl. Lorem", Pekerjaan: "Backend", Alasan: "Alasan Fitri"},
}

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/login", loginHandler)
	http.ListenAndServe(":8080", nil)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	email := r.URL.Query().Get("email")
	user, found := users[email]
	if !found {
		http.Error(w, "User not found", http.StatusNotFound)
		return
	}

	tmpl := template.Must(template.ParseFiles("index.html"))
	err := tmpl.Execute(w, user)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func loginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodPost {
		email := r.FormValue("email")
		_, found := users[email]
		if found {
			http.Redirect(w, r, fmt.Sprintf("/?email=%s", email), http.StatusSeeOther)
			return
		}
	}

	tmpl := template.Must(template.ParseFiles("login.html"))
	err := tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}
