package main

import "time"

type Medico struct {
	ID      int
	Nome    string
	Funcao  string
}
	
type Consulta struct {
	MedicoNome string
	Funcao     string
	Data       time.Time
}

type DashboardData struct {
	Especialidades          []string
	EspecialidadeSelecionada string
	Medicos                 []Medico
	Consultas               []Consulta
}