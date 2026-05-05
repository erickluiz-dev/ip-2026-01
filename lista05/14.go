// 14. Faça um programa que leia dois vetores de 10 elementos inteiros cada um e mostre o vetor resultante da
// intercalação desses dois vetores.
package main

import f "fmt"

func main(){
	var array1 [10]int
	var array2 [10]int

	f.Println("---Vetor 1---")

	for i := 0; i < len(array1); i++{
		f.Scan(&array1[i])
	}

	f.Println("---Vertor 2---")

	for i := 0; i < len(array2); i++{
		f.Scan(&array2[i])
	}
	var result_arr []int

	for i := 0; i < len(array1) || i < len(array2); i++{
		result_arr = append(result_arr, array1[i])
		result_arr = append(result_arr, array2[i])
	} 

	f.Println("---Vetor Resultante---")
	f.Print(result_arr)
}