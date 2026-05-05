// 10. A série de Fibonacci é formada pela sequência: 1, 1, 2, 3, 5, 8, 13, 21, 34, 55, ... Escreva um programa que armazene
//  em um vetor os primeiros 50 termos da série de Fibonacci. Após isso, o programa deve imprimir todos os valores armazenados.

package main

import f "fmt"

func main(){
	x, y := 1, 0
	var arr [50]int
	for i := 0; i < len(arr); i++{
		x, y = x + y, x
		arr[i] = x
	}
	for _, v := range arr{
		f.Printf("%d ", v)
	}
}