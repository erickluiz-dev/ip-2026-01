// 18. Uma locadora tem as seguintes regras para aluguel de DVDs:
// - Às segundas, terças e quintas (2, 3 e 5) : desconto de 40% em relação ao preço normal;
// - Às quartas , sextas, sábados e domingos (4,6 ,7 e 1): preço normal;
// - Aluguel de DVDs comuns: preço normal;
// - Aluguel de lançamentos: acréscimo de 15% em relação ao preço normal.
// Desenvolver um programa para ler o preço normal do DVD alugado (em R$) e sua categoria (comum ou
// lançamento). Calcular e imprimir o preço final que será pago pela locação do DVD.
package main

import f "fmt"

func main() {
	day, price, releasestr, release := 0, 0.0, "nao", false
	f.Print("Qual é o dia da semana? ")
	f.Scan(&day)
	f.Print("Qual é o preço do dvd? ")
	f.Scan(&price)
	f.Print("O DVD é um lançamento? ")
	f.Scan(&releasestr)
	if releasestr == "sim"{
		release = true 
	}
	b := buy(day, price, release)
	if day >= 1 && day <= 8{
	f.Printf("O preço pela locaçao do DVD é R$ %.2f ", b)
	}else{
		f.Print("Dia da semana invalido ")
	}
}
func buy(number1 int, number2 float64, word bool) float64{
	end_buy := number2
	if number1 == 2|| number1 == 3|| number1 == 5{
		end_buy *= 0.6
	}
	if word {
		end_buy *= 1.15
	}
	return end_buy 
}