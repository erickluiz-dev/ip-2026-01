package main

import (
	"html/template"
	"net/http"
)

func DashboardHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	tmpl, err := template.ParseFiles("frontend/dashboard-usuario.html")
	if err != nil {
		http.Error(w, "Erro ao carregar o painel do paciente: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}