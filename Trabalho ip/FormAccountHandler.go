package main

import(
	"net/http"
	"html/template"

)

func FormAccountHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("frontend/criar-conta-usuario.html")
		if err != nil {
			http.Error(w, "Erro ao abrir página: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		email := r.FormValue("email")
		senha := r.FormValue("senha")
		nome := r.FormValue("nome")
		cpf := r.FormValue("cpf")

		query := "INSERT INTO usuarios (email, senha, nome, cpf) VALUES ($1, $2, $3, $4)"

		_, err := db.Exec(query, email, senha, nome, cpf)
		if err != nil {
			http.Error(w, "Erro ao salvar no Banco de Dados: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}