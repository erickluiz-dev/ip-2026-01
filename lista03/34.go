// 34. Faça um programa que leia dois números inteiros positivos N1 e N2, calcule e escreva o mínimo múltiplo
// comum para este par de números (N1,N2).
package main

import f "fmt"

func main(){
	n1, n2 := 0, 0
	f.Print("Digite o 1º número ")
	f.Scan(&n1)
	f.Print("Digite o 2º número ")
	f.Scan(&n2)

	x := mmc(n1, n2)
	f.Printf("O MMC de %d e %d é %d ", n1, n2, x)
}
func mmc(num1, num2 int) int{
	x1, x2 := 0, 0 
	if num1 == num2{
		return num1
	}
	for{
		if x1 < x2 {
			x1 += num1
		}else{
			x2 += num2
		}

		if x1 == x2{
			return x1
		}
	}
}