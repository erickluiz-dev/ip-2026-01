// Faça um programa que receba 100 valores numéricos, armazene-os em um vetor, calcule e imprima o valor do somatório dado a seguir:
// S = (b0-b99)^3 + (b1-b98)^3 + (b2-b97)^3 + . . . + (b49-b50)^3
package main

import (
	f "fmt"
	m "math"
)

func main() {
	var arr [4]float64
	for i := 0; i < len(arr); i++ {
		f.Scan(&arr[i])
	}
	S, j := 0.0, len(arr)-1
	for i := 0; i < len(arr)/2; i++{
		S += m.Pow((arr[i] - arr[j]), 3)
		j--
	}
	f.Printf("%.2f ", S)
}