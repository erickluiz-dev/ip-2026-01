// 29. Escreva um programa que calcule e imprima o valor do somatório de todos os números inteiros de 1 a N, onde
// N é um número inteiro positivo fornecido pelo usuário.
package main

import f "fmt"

func main(){
	N, S := 0, 0
	f.Print("Digite um valor para N \n")
	f.Scan(&N)

	for i:=1; i <= N; i++{
		S += i
	} 
	f.Printf("A soma é = %d ", S)
}