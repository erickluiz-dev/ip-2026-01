// 9. Faça um programa que receba 2 notas de N alunos. Calcule e mostre:
// a) a média aritmética das 2 notas de cada aluno;
// b) uma mensagem de acordo com as regras a seguir:
// Média Aritmética Mensagem
// Até 3 Reprovado
// Entre 3 e 7 Exame
// De 7 para cima Aprovado
// c) o total de alunos aprovados;
// d) o total de alunos de exame;
// e) o total de alunos reprovados;
// f) a média da classe.
// Assuma que o N informado é válido, assim como as 2 notas de cada aluno.
package main

import f "fmt"

type Dados struct {
	N1       float64
	N2       float64
	Situação string
}

func main() {
	N := 0
	f.Print("Digite o número de alunos: ")
	f.Scan(&N)

	alunos := make([]Dados, N)
	totalAprovados, totalExame, totalReprovados := 0, 0, 0
	somaMediasClasse := 0.0

	for i := 0; i < N; i++ {
		f.Printf("\n--- Aluno %d ---\n", i+1)
		f.Print("Digite a 1ª nota: ")
		f.Scan(&alunos[i].N1)

		f.Print("Digite a 2ª nota: ")
		f.Scan(&alunos[i].N2)

		media := (alunos[i].N1 + alunos[i].N2) / 2
		somaMediasClasse += media

		if media <= 3 {
			alunos[i].Situação = "Reprovado"
			totalReprovados++
		} else if media < 7 {
			alunos[i].Situação = "Exame"
			totalExame++
		} else {
			alunos[i].Situação = "Aprovado"
			totalAprovados++
		}

		f.Printf("Média: %.2f | Situação: %s\n", media, alunos[i].Situação)
	}

	mediaClasse := somaMediasClasse / float64(N)

	f.Println("\n--- RESULTADOS FINAIS ---")
	f.Printf("c) Total Aprovados: %d\n", totalAprovados)
	f.Printf("d) Total Exame: %d\n", totalExame)
	f.Printf("e) Total Reprovados: %d\n", totalReprovados)
	f.Printf("f) Média da Classe: %.2f\n", mediaClasse)
}