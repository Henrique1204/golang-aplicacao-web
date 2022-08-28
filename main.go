package main

import (
	"net/http"
	"text/template"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var tmp = template.Must(template.ParseGlob("templates/*.html"))

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camisa", Descricao: "Camisa vermelha ou azul", Preco: 20.50, Quantidade: 2},
		{Nome: "Short", Descricao: "Short verde ou amarelo", Preco: 39.99, Quantidade: 4},
		{Nome: "Fone", Descricao: "Fone com dois lados funcionando", Preco: 50, Quantidade: 1},
	}

	tmp.ExecuteTemplate(w, "Index", produtos)
}

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}
