// Escreva um programa que armazene os 100 primeiros números ímpares em um vetor. Imprima o vetor em
// seguida.
package main

import f "fmt"

func main(){
	j := 0
	var arr []int
	for{
		if j % 2 != 0{
			arr = append(arr, j)
		}
		if len(arr) == 100{
			break
		}
		j++
	}
	for _, v := range arr{
		f.Printf("%d ", v)
	}
}