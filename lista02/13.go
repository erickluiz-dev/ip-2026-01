// 13. Escreva um programa que receba um número inteiro positivo de 3 casas e imprima o algarismo da casa das
// dezenas. Não se esqueça de testar para ver se o número informado tem realmente 3 casas
package main

import f "fmt"	

func main(){
	var numero string
	f.Print("Digite um número com 3 casas ")
	f.Scan(&numero)
	l := len(numero)
	c := numero[1]
	if l == 3 && c >= '0' && c <= '9'{
	f.Printf("O algarismo da casa das dezenas é %c ", c)
	}else{
		f.Print("Número invalido ")
	}
}