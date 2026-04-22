// 37. Faça um programa que leia um número inteiro positivo na base 8, calcule e imprima o seu equivalente na base 10.
package main

import (
	f "fmt"
	m "math"
)

func main() {
	x := 0
	f.Print("Digite um número ")
	f.Scan(&x)

	y := decompor(x)
	f.Print(y)
}

func decompor(x int) float64{
	var arr []float64

	for x > 0 {
		arr = append(arr, float64(x % 10))
		x /= 10
	}
	return tradução(arr)
}

func tradução(arr []float64) float64 {
	x := 0.0
	for i := 0; i < len(arr); i++{
		x += arr[i] * m.Pow(8, float64(i))
	}
	return x
}