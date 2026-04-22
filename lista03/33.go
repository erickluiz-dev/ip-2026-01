// 33. Faça um programa que: 
// - leia dois números inteiros positivos (N1 e N2);
// - calcule e escreva para este par de números (N1 e N2), o quociente e o resto da divisão de N1 por N2. Obs.: a
// máquina que irá calcular o quociente e o resto desta divisão
package main

import f "fmt"

func main(){
	n1, n2 := 0, 0
	f.Print("Digite o divisor ")
	f.Scan(&n1)
	f.Print("Digite o dividendo ")
	f.Scan(&n2)

	div, res := divisao(n1, n2)
	f.Printf("%d / %d = %d resto = %d ", n1, n2, div, res)
}
func divisao(num1, num2 int) (int, int){
	x:=0 
	for{
		if num1< num2{
			return x, num1
		}
		num1 -= num2
		x++

	}
}