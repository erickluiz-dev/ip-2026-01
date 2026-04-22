// 4. Escreva um programa que receba vários números inteiros e verifique se eles são ou não quadrados perfeitos. O // programa deve terminar 
// quando o usuário informar um número menor ou igual a zero. Obs.: Um número é quadrado perfeito quando tem um número inteiro como raiz quadrada.
// Não é permitido usar o comando sqrt em sua solução.
package main

import f "fmt"

func main() {
	number, test := 0, true
	
	for{
	f.Print("Digite um número ")
	f.Scan(&number)
	test = true

	if number <= 0{
		return
	}

	for i:= 1; i <= number; i++{
		if(i * i) == number{
			f.Printf("O número %d é um quadrado perfeito \n\n", number)
			test = false
			break	
		}
	}
	
	if test{
		f.Printf("O número %d não é um Quadrado Perfeito \n\n", number)
	}
}
}