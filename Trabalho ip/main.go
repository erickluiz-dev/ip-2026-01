package main

import (
	"database/sql"
	f "fmt"
	"net/http"
	"os/exec"
	"runtime"
	"time"

	_ "github.com/lib/pq"
)

var db *sql.DB

func abrirNavegador(url string) {
	var cmd *exec.Cmd

	switch runtime.GOOS {
	case "windows":
		cmd = exec.Command("cmd", "/c", "start", url)
	case "darwin": // Mac
		cmd = exec.Command("open", url)
	case "linux":
		cmd = exec.Command("xdg-open", url)
	default:
		return
	}

	_ = cmd.Start()
}

func main() {
	var err error

	conectarStr := "user=postgres password=esufg123 dbname=clinica_ip host=localhost port=5432 sslmode=disable"

	db, err = sql.Open("postgres", conectarStr)
	if err != nil {
		panic("Erro ao configurar o banco de dados: " + err.Error())
	}

	err = db.Ping()
	if err != nil {
		panic("Não foi possível conectar ao banco de dados: " + err.Error())
	}

	f.Println("✅ Conectado ao banco de dados PostgreSQL com sucesso!")

	http.HandleFunc("/FormAccount", FormAccountHandler)
	http.HandleFunc("/FormMedical", FormMedicalHandler)
	http.HandleFunc("/", IndexHandler)
	http.HandleFunc("/Login", LoginHandler)
	http.HandleFunc("/Dashboard", DashboardHandler)

	porta := ":5500"
	url := "http://localhost" + porta
	f.Println("🚀 Servidor rodando na porta", porta)

	go func() {
		time.Sleep(500 * time.Millisecond)
		f.Println("🌐 Abrindo o navegador em:", url)
		abrirNavegador(url)
	}()

	err = http.ListenAndServe(porta, nil)
	if err != nil {
		panic("Erro ao iniciar o servidor: " + err.Error())
	}
}