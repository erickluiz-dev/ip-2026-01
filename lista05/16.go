// 16. Dado um vetor com dados de 50 idades, elabore um programa que permita calcular a moda das idades. Obs.:
// Moda é o valor que tem maior incidência de repetições.
package main

import f "fmt"

func main(){
	ages := map[int]int{}
	var arr [50]int
	moda := 0
	repeticoes := 0

	for i := 0; i < len(arr); i++{
		f.Scan(&arr[i])
		ages[arr[i]]++

		if repeticoes < ages[arr[i]]{
			moda = arr[i]
			repeticoes = ages[arr[i]]
		}
	}
	f.Printf("Moda: %d ", moda)
}