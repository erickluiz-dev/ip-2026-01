package main

import (
	"database/sql"
	f "fmt"
	"net/http"
)

var db *sql.DB

func main() {

	var err error

	conectarStr := "user=postgres password=esufg123 dbname=Usuário_Clinica_IP host=localhost port=5432 sslmode=disable"

	db, err := sql.Open("postgrer", conectarStr)
	if err != nil{
		panic("Erro ao configurar o banco de dados: " + err.Error())
	}

	defer db.Close()

	err = db.Ping()
	if err != nil{
		panic("Não foi possível conectar ao banco de dados: " + err.Error())
	}
	
	f.Println("✅ Conectado ao banco de dados PostgreSQL com sucesso!")

	http.HandleFunc("/FormAccount", FormAccountHandler)
	http.HandleFunc("/FormMedical", FormMedicalHanndler)

	porta := "5500"
	f.Println("🚀 Servidor rodando na porta", porta)

	err = http.ListenAndServe(porta, nil)
	if err != nil {
		panic("Erro ao iniciar o servidor: " + err.Error())
	}
}	