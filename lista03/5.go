// 5. Escreva um programa que receba a idade, a altura e o peso de várias pessoas. Calcule e imprima:
// - a quantidade de pessoas com idade superior a 50 anos;
// - a média das alturas das pessoas com idade entre 10 e 20 anos;
// - a porcentagem de pessoas com peso inferior a 40 quilos entre todas as pessoas analisadas.
// Considere que os dados informados são válidos. Pergunte ao usuário se ele deseja continuar digitando dados ou
// não (Exemplo: 1 - Sim, Outro valor diferente de 1 - Não).
package main

import (
	f "fmt"
)

func main(){
	weight := 0.0
	height := 0.0
	age := 0.0
	count_age := 0
	count_height := 0.0
	count_weight := 0.0
	sum_height := 0.0
	count_people := 0.0
	end := 1

	for{
		end  = 0
		f.Print("Digite a idade ")
		f.Scan(&age)
		f.Print("Digite a altura ")
		f.Scan(&height)	
		f.Print("Digite o peso ")	
		f.Scan(&weight)


		if age > 50{
			count_age++
		}

		if age >= 10 && age <= 20{
			sum_height += height
			count_height++
		}

		if weight < 40{
			count_weight++
		}

		count_people++

		f.Print("Deseja continuar? ")
		f.Scan(&end)
		if end != 1{
			break
		}
	}
	percentage := (100 * count_weight) / count_people
	average := (sum_height / count_height)
	f.Printf("%d pessoas tem mais de 50 anos \n%.2f é a média das altura de pessoas com mais de 10 e menos 20 \nA porcentagem de porcentagem de pessoas com menos de 40 kg é %.0f%", count_age, average, percentage)
}