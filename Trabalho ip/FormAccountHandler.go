package main

import(
	f "fmt"
	"database/sql"
	"net/http"

)

var db *sql.DB

func creatAccountHandler(w http.ResponseWriter, r http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	senha := r.FormValue("senha")
	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")

	query := "INSERT INTO usuarios (email, senha, nome, cpf) VALUES ($1, $2, $3, $4)"

	_, err := db.Exec(query, email, senha, nome, cpf)
	if err != nil{
		http.Error(w, "Erro ao salvar no Banco de Dados"+err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	f.Fprintln(w, "<h2>Conta criada com sucesso!</p>")

}