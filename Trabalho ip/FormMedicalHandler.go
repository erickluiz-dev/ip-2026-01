package main

import(
	"net/http"
	"html/template"
)

func FormMedicalHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method == http.MethodGet {

		tmpl, err := template.ParseFiles("frontend/criar-conta-medico.html")
		if err != nil {
			http.Error(w, "Erro ao abrir página: "+err.Error(), http.StatusInternalServerError)
			return
		}

		tmpl.Execute(w, nil)
		return
	}

	if r.Method == http.MethodPost {

		nome := r.FormValue("nome")
		email := r.FormValue("email")
		senha := r.FormValue("senha")
		cpf := r.FormValue("cpf")
		atuacao := r.FormValue("atuacao")

		query := "INSERT INTO medico (nome, email, senha, cpf, funcao) VALUES ($1, $2, $3, $4, $5)"

		_, err := db.Exec(query, nome, email, senha, cpf, atuacao)
		if err != nil {
			http.Error(w, "Erro ao salvar no banco de dados: "+err.Error(), http.StatusInternalServerError)
			return
		}

		http.Redirect(w, r, "/", http.StatusSeeOther)
		return
	}

	http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
}