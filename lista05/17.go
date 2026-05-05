// 17. Escreva um programa que leia um vetor de dez elementos inteiros e mostre os números primos e suas respectivas posições.
package main

import f "fmt"

func main(){
	var arr [5]int

	for i := 0; i < len(arr); i++{
		f.Scan(&arr[i])
	}

	f.Println("---Números Primos---")

	teste := 0

	for i, v := range arr{
		if primo(v){
			f.Printf("Númeoro: %d | Posição: %d \n", v, i)
			teste++
		}
	}
}

func primo(num int) bool{
	x := 0
	for i := 1; i <= (num/2); i++{
		if num % i == 0 {
			x++
		}
	}
	if x == 1 {
		return true
	}else{
		return false
	}
}