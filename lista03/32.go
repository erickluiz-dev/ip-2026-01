// 32. Faça um programa que calcule e escreva a multiplicação de dois números N1 e N2 lidos do teclado. Obs.: a
// máquina que irá executar esse programa somente sabe adicionar e subtrair.
package main

import f "fmt"

func main(){
	n1, n2 := 0, 0
	f.Print("Digte um valor para N1 \n")
	f.Scan(&n1)
	f.Print("Digite um valor para N2 \n")
	f.Scan(&n2)

	res := n1

	for i:= 1; i < n2; i++{
		res += n1
	}
	f.Printf("%d x %d = %d ", n1, n2, res)
}