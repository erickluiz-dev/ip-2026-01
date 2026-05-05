// 13. Uma grande firma deseja saber quais os três empregados mais recentes. Faça um programa para ler um número indeterminado de 
// informações (máximo de 100), contendo o número do empregado e o número de meses de trabalho deste empregado. Imprima os três 
// empregados que entraram para trabalhar mais recentemente na firma. Obs.: a última informação contém os dois números iguais a zero. 
// Assuma que não existem dois ou mais empregados admitidos no mesmo mês.
package main

import f "fmt"

func main(){
	type Funcionario struct{
		Num int
		Meses int
	}
	

	funcionarios_novos := [3]Funcionario{
		{0, 9999}, {0, 9999}, {0, 9999},
	}

	for i := 0; i < 100; i++{
		f.Print("Funcionário: ")
		id := 0
		f.Scan(&id)
		f.Print("Meses: ")
		meses := 0
		f.Scan(&meses)
		
		if id == 0 && meses == 0 {break}

		if funcionarios_novos[0].Meses > meses{
			funcionarios_novos[2] = funcionarios_novos[1]
			funcionarios_novos[1] = funcionarios_novos[0]
			funcionarios_novos[0] = Funcionario{id, meses}
		}else if meses < funcionarios_novos[1].Meses{
			funcionarios_novos[2] = funcionarios_novos[1]
			funcionarios_novos[1] = Funcionario{id, meses}
		}else if meses < funcionarios_novos[2].Meses{
			funcionarios_novos[2] = Funcionario{id,meses}
		}
	}

	f.Print("---Funcionários Mais Recentes---\n")

	for _, v := range funcionarios_novos{
		f.Printf("Funcionário: %d | Meses: %d \n", v.Num, v.Meses)
	}
}