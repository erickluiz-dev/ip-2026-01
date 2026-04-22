// 10. Faça um programa que monte a sequência de Fibonacci com N termos.
// A sequência de Fibonacci é dada por: 0 – 1 – 1 – 2 – 3 – 5 – 8 – 13 – 21 – 34 – 55 – ...
// O primeiro e o segundo termos da sequência de Fibonacci são 0 e 1.
// Considere que o usuário irá informar um número N >= 3.
package main

import f "fmt"

func main(){
	n1 := 0
	n2 := 1
	n := (n1 + n2)

	N := 0
	f.Print("Digite N ")
	f.Scan(&N)

	f.Printf("%d - %d - %d ", n1, n2, n)

	for i := 0; i < (N - 3); i++{

		n1 = n
		n = n2 + n
		n2 = n1

		f.Printf("- %d ", n)
	}
}