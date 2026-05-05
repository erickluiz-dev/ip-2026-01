// 6. Escreva um programa que armazene todos os números inteiros de 100 a 1 (ordem decrescente) em um vetor. A
// seguir, imprima os elementos do vetor.
package main

import f "fmt"

func main(){
	var arr []int

	for i := 100; i > 0; i--{
		arr = append(arr, i)
	}
	for _, v := range arr{
		f.Printf("%d ", v)
	}
}