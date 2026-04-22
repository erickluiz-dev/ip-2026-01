// 15. Seja a seguinte série: 1, 4, 9, 16, 25, 36, ...Escreva um programa que gere esta série até o N-ésimo termo. N será informado pelo usuário.
package main

import f "fmt"

func main(){
	N, x, s :=  0, 1, 0
	f.Print("Digirte N ")
	f.Scan(&N)
	for i:=0; i<N; i++{
		s += x
		f.Printf("%d ", s)
		x += 2
	}
}