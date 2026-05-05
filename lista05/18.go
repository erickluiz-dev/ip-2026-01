// Faça um programa que leia um vetor de elementos inteiros de 10 posições, mas já o leia de maneira ordenada
// crescente. Ao final, imprima o vetor.
package main

import f "fmt"

func main(){
	var array [10]int
	for i := 0; i < len(array); i++{
		x:= 0
		f.Scan(&x)
		j := i-1
		for j >= 0 && array[j] > x{
			array[j+1] = array[j]
			j--
		}
		array[j+1] = x
	}
	f.Print(array)
}