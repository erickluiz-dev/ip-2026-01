package main

import(
	f "fmt"
	"net/http"
	"html/template"

)

func FormAccountHandler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("/frontend/criar-conta-usuario.html")
	if err != nil {
		http.Error(w, "Erro ao abrir página: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	email := r.FormValue("email")
	senha := r.FormValue("senha")
	nome := r.FormValue("nome")
	cpf := r.FormValue("cpf")

	query := "INSERT INTO usuarios (email, senha, nome, cpf) VALUES ($1, $2, $3, $4)"

	_, err = db.Exec(query, email, senha, nome, cpf)
	if err != nil{
		http.Error(w, "Erro ao salvar no Banco de Dados"+err.Error(), http.StatusInternalServerError)
		return
	}
	
	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	f.Fprintln(w, "<h2>Conta criada com sucesso!</p>")

}