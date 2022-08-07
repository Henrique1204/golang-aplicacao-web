package main

import (
	"net/http"
	"text/template"
)

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	tmp.ExecuteTemplate(w, "Index", nil)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
