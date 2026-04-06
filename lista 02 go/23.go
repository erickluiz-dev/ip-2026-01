// 23. Criar um programa que leia a idade de uma pessoa e que mostre a sua classe eleitoral:
// - Não-eleitor (abaixo de 16 anos);
// - Eleitor obrigatório (entre 18 e 65 anos);
// - Eleitor facultativo (entre 16 e 18 anos e maior de 65 anos).
package main

import f "fmt"

func main(){
	idade := 0
	f.Print("Qual é a sua idade? ")
	f.Scan(&idade)
	classe := "Eleitor facultativo"
	if idade < 16{
		classe = "Não-eleitor"
	}else if idade >= 18 && idade <=65{
		classe = "Eleitor obrigatório"
	}
	f.Printf("Sua classe eleitoral é %s ", classe)
}