package main

import(
	f "fmt"
	"net/http"
	"html/template"
)

func FormMedicalHanndler(w http.ResponseWriter, r *http.Request) {

	tmpl, err := template.ParseFiles("/frontend/criar-conta-medico.html")
	if err != nil {
		http.Error(w, "Erro ao abrir página: "+err.Error(), http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)

	if r.Method != http.MethodPost{
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	nome := r.FormValue("nome")
	email := r.FormValue("email")
	senha := r.FormValue("senha")
	cpf := r.FormValue("cpf")
	atuacao := r.FormValue("atuacao")

	quary := "INSERT INTO medico (nome, email, senha, cpf, funcao) VALUES ($1, $2, $3, $4, $5)"

	_, err = db.Exec(quary, nome, email, senha, cpf, atuacao)
	if err != nil{
		http.Error(w, "Erro ao salvar no banco de dados"+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	f.Fprintln(w, "<h2>Conta médica criada com sucesso!</h2>")
}