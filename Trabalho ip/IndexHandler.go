package main

import (
	"html/template"
	"net/http"
)

func IndexHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("frontend/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar index: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}