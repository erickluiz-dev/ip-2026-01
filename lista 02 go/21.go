// 21. Escrever um programa que leia o número de identificação, as 3 notas obtidas por um aluno nas 3 verificações e
// a média dos exercícios que fazem parte da avaliação. Calcular a média de aproveitamento do aluno, usando a
// fórmula:
// Média Final =

// ( nota1 + nota2 ∗ 2 + nota3 ∗ 3 + médiados exercícios )

// 7

// e o seu conceito, usando a tabela a seguir:

// Média de Aproveitamento Conceito
// >= 9,0 e <= 10,0 A
// >=7,5 e < 9,0 B
// >=6,0 e < 7,5 C
// >=4,0 e < 6,0 D
// < 4,0 E

// O programa deve escrever o número do aluno, suas notas, a média dos exercícios, a média de aproveitamento,
// o conceito correspondente e a mensagem: APROVADO se o conceito for A, B ou C e REPROVADO, se o
// conceito for D ou E.

package main

import f "fmt"

func main(){
	var nome string
	f.Print("Qual é o nome do aluno? ")
	f.Scan(&nome)
	const NumNotas int = 4
	var (
		nota [NumNotas]float64 
	)

	for i := 0; i < NumNotas; i++{
		if i == 3{
			f.Printf("Digite a media dos exercícios ")
		}else{
		f.Printf("Digite a %dº nota ", i+1)
		}
		f.Scan(&nota[i])
	}
	m := media(nota)
	c, a := conceito(m)
	
	f.Printf("Aluno : %s \n", nome)
	
	for i, n := range nota{
		if i == 3{
			f.Printf("Média dos execícios: %.2f \n", n)
		}else{
			f.Printf("A %dº nota é %.2f \n", i+1, n)
		}
	}
	f.Printf("Média de aproveitamento : %.2f \n", m)
	f.Printf("%s %s \n", c, a)
}
func media(n [4]float64) float64{
calculo := (n[0] + (2 * n[1]) + (3 * n[2]) +  n[3]) / 7
return calculo 
}

func conceito(number float64) (string, string){
	con, apro := "", "Reprovado"
	if number >= 9.0 && number <= 10.0{
		con = "A"
	}else if number >=7.5 && number < 9.0{
		con = "B"
	}else if number >=6.0 && number < 7.5{
		con = "C"
	}else if number >=4.0 && number < 6.0{
		con = "D"
	}else{
		con = "E"
	}
	if con == "A" || con == "B" || con == "C"{
		apro = "Aprovado"
	}
	return con, apro
}