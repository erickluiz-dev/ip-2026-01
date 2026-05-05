// 9. Escreva um programa que receba a altura de 10 atletas. Esse programa deve imprimir as alturas daqueles
// atletas que têm altura maior do que a média.
package main

import f "fmt"

func main(){
	var arr [10]float64
	sum := 0.0
	for i := 0; i < len(arr); i++{
		f.Scan(&arr[i])
		sum += arr[i]
	}
	average := sum / float64(len(arr))
	for _, v := range arr{
		if v > average{
			f.Printf("%f ", v)
		}
	}
}