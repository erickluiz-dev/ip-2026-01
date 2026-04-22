// 14. Faça um programa que imprima todos os números primos existentes entre N1 e N2, em que N1 e N2 são números naturais fornecidos pelo
// usuário.
package main

import f "fmt"

func main(){
	N1, N2 := 0, 0
	f.Print("Digite N1 e N2 ")
	f.Scan(&N1)
	f.Scan(&N2)


	for i := N1; i <= N2; i++{
		r := primo(i)
		if r {
			f.Printf("%d ", i)
		}
	}
}
func primo(num int) bool{
	if num < 2{
		return false
	}
	for i := 2; i < num; i++{
		if num % i == 0{
			return false	
		}
	}
		return true
}