package main

import f "fmt"

func main(){
	var arr [10]int
	for i := 0; i < len(arr); i++{
		f.Scan(&arr[i])
	}
	menor, indice := arr[0], 0
	for i, v := range arr{
		if v < menor{
			menor = v
			indice = i
		}
	}
	f.Printf("O menor elemento do vetor é %d e sua posição dentro do vetor é: %d", menor, indice+1) 
}