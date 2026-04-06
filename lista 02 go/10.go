package main

import(
	 f "fmt"
	 "strings"
)

func main(){
	v := false
	x := 0
	f.Print("Qual é o  local da sua viagem \n1 - Região Norte \n1 - Região Norte \n3 - Região Centro-Oeste \n4 - Região Sul\n")
	f.Print("Digite o número de 1-4 de acordo com locar da sua viagem ")
	f.Scan(&x)
	var volta string
	f.Print("Sua viagem é de ida e volta ")
	f.Scan(&volta)
	volta = strings.Trim(volta, " ")
	if volta == "sim" || volta == "Sim"{
		v = true
	}
	p := preco(x, v)
	f.Printf("O preço da viagem é %.2f ",p )
}
func preco(numero int, volta bool) float64 {
	valor := 0.0
	if volta {
		if numero == 1{
		valor = 900
	}else if numero == 2{
		valor = 650
	}else if numero == 3{
		valor = 600
	}else if numero == 4{
		valor = 550
	}}else{
	if numero == 1{
		valor = 500
	}else if numero == 2{
		valor = 350
	}else if numero == 3{
		valor = 350
	}else if numero == 4{
		valor = 300
	}
}
return valor
}