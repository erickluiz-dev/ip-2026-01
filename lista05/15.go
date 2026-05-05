// 15. Crie um programa que leia um vetor de 30 números inteiros e gere um segundo vetor cujas posições pares
// possuirão elementos que serão o dobro do vetor original e as ímpares, o triplo.
package main

import f "fmt"

func main(){
	var arr [4]int

	f.Println("---Primeiro Vetor---")

	for i := 0; i < len(arr); i++{
		f.Scan(&arr[i])
	}
	arr2 := make([]int, len(arr))
	for i := 0; i < len(arr); i++{
		if i % 2 == 0 {
			arr2[i] = 2 * arr[i]
		}else{
			arr2[i] = arr[i] * 3
		}
	}

	f.Println("---Segundo Vetor---")
	f.Print(arr2)
}