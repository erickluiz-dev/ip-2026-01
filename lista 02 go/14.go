// 14. Desenvolva um programa para calcular e imprimir o preço final de um carro. O valor do preço inicial de
// fábrica é fornecido pelo usuário. O carro pode ter as seguintes opções:
// a) Ar condicionado (R$ 1750,00)
// b) Pintura metálica (R$ 800,00)
// c) Vidro elétrico (R$ 1200,00)
// d) Direção hidráulica (R$ 2000,00)
package main

import f "fmt"

func main(){
	preco := 0.0
	f.Print("Digite o preço inicial ")
	f.Scan(&preco)

	var ar string
	arc := false
	f.Print("Você deseja ar condicionado? ")
	f.Scan(&ar)
	if ar == "s"{
		arc = true
	}

	var pi string
	pim := false
	f.Print("Você deseja pintura metálica? ")
	f.Scan(&pi)
	if ar == "s"{
		pim = true
	}

	var vi string
	vie := false
	f.Print("Você deseja vidro elétrico? ")
	f.Scan(&vi)
	if ar == "s"{
		vie = true
	}	

	var di string
	dih := false
	f.Print("Você deseja direção hidráulica? ")
	f.Scan(&di)
	if ar == "s"{
		dih = true
	}
p := precofinal(preco, arc, pim, vie, dih)
f.Print("O preço final do carro é R$ %.2f ", &p)
}
func precofinal(preco float64, arc, pim, vie, dih bool) float64{
	precof := preco
	if arc {
		 precof += 1750
	}
	if pim{
		precof += 800
	}
	if vie{
		precof += 1200
	}
	if dih {
		precof += 2000
	}
	return precof
}