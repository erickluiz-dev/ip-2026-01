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

	especialidade := r.URL.Query().Get("atuacao")

	// =========================
	// BUSCAR ESPECIALIZAÇÕES
	// =========================

	queryEspecialidades := `
		SELECT DISTINCT funcao
		FROM medicos
	`

	rowsEspecialidades, err := db.Query(queryEspecialidades)
	if err != nil {
		http.Error(w, "Erro ao buscar especializações: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rowsEspecialidades.Close()

	var especialidades []string

	for rowsEspecialidades.Next() {

		var funcao string

		err := rowsEspecialidades.Scan(&funcao)
		if err != nil {
			http.Error(w, "Erro ao ler especializações: "+err.Error(), http.StatusInternalServerError)
			return
		}

		especialidades = append(especialidades, funcao)
	}

	// =========================
	// BUSCAR MÉDICOS FILTRADOS
	// =========================

	var medicos []Medico

	if especialidade != "" {

		queryMedicos := `
			SELECT id, nome, funcao
			FROM medicos
			WHERE funcao = $1
		`

		rowsMedicos, err := db.Query(queryMedicos, especialidade)
		if err != nil {
			http.Error(w, "Erro ao buscar médicos: "+err.Error(), http.StatusInternalServerError)
			return
		}

		defer rowsMedicos.Close()

		for rowsMedicos.Next() {

			var medico Medico

			err := rowsMedicos.Scan(
				&medico.ID,
				&medico.Nome,
				&medico.Funcao,
			)

			if err != nil {
				http.Error(w, "Erro ao ler médicos: "+err.Error(), http.StatusInternalServerError)
				return
			}

			medicos = append(medicos, medico)
		}
	}

	// =========================
	// BUSCAR CONSULTAS
	// =========================

	queryConsultas := `
		SELECT medicos.nome, medicos.funcao, consultas.data_consulta
		FROM consultas
		INNER JOIN medicos
		ON consultas.medico_id = medicos.id
		WHERE consultas.paciente_id = 1
	`

	rowsConsultas, err := db.Query(queryConsultas)
	if err != nil {
		http.Error(w, "Erro ao buscar consultas: "+err.Error(), http.StatusInternalServerError)
		return
	}

	defer rowsConsultas.Close()

	var consultas []Consulta

	for rowsConsultas.Next() {

		var consulta Consulta

		err := rowsConsultas.Scan(
			&consulta.MedicoNome,
			&consulta.Funcao,
			&consulta.Data,
		)

		if err != nil {
			http.Error(w, "Erro ao ler consultas: "+err.Error(), http.StatusInternalServerError)
			return
		}

		consultas = append(consultas, consulta)
	}

	data := DashboardData{
		Especialidades:           especialidades,
		EspecialidadeSelecionada: especialidade,
		Medicos:                  medicos,
		Consultas:                consultas,
	}

	tmpl, err := template.ParseFiles("frontend/dashboard-usuario.html")
	if err != nil {
		http.Error(w, "Erro ao carregar dashboard: "+err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, data)
}