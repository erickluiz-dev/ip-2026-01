package main

import (
	"net/http"
)

func AgendarConsultaHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	medicoID := r.FormValue("medico")
	atuacao := r.FormValue("atuacao")
	dataConsulta := r.FormValue("data")


	pacienteID := 1

	query := `
		INSERT INTO consultas (paciente_id, medico_id, data_consulta)
		VALUES ($1, $2, $3)
	`

	_, err := db.Exec(query, pacienteID, medicoID, dataConsulta)
	if err != nil {
		http.Error(w, "Erro ao salvar consulta: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "text/html; charset=utf-8")
	http.Redirect(
		w,
		r,
		"/Dashboard?atuacao="+atuacao,
		http.StatusSeeOther,
)
}