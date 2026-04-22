// 11. Escreva um programa que calcule o fatorial de um número inteiro N fornecido pelo usuário. Cuidado com valores inválidos!
package main

import f "fmt"

func main(){
	N := 0
	f.Print("Digite um número ")
	f.Scan(&N)

	r := fat(N)
	f.Printf("%d ", r)
}
func fat(n int) int{
	if n == 0{
		return 1
	}
	return n * fat(n - 1)
}