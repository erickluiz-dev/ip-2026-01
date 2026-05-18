package main

import (
	"database/sql"
	"net/http"
)

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido. Use POST para logar.", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	senha := r.FormValue("senha")

	var usuarioNome string

	query := "SELECT nome FROM usuarios WHERE email = $1 AND senha = $2"
	err := db.QueryRow(query, email, senha).Scan(&usuarioNome)

	if err != nil {
		if err == sql.ErrNoRows {
			w.Header().Set("Content-Type", "text/html; charset=utf-8")
			w.WriteHeader(http.StatusUnauthorized)
			w.Write([]byte(`<h3>E-mail ou senha incorretos!</h3><br><a href="/">Voltar e tentar novamente</a>`))
			return
		}
		http.Error(w, "Erro interno no servidor: "+err.Error(), http.StatusInternalServerError)
		return
	}


	http.Redirect(w, r, "/Dashboard", http.StatusSeeOther)
}