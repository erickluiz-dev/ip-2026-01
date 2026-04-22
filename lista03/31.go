// 31. Faça um programa que calcule e escreva o número de grãos de milho que se pode colocar em um tabuleiro e xadrez, colocando 1 
// no primeiro quadro e nos quadros seguintes o dobro do quadro anterior. São 64 quadros no total.
package main

import f "fmt"

func main(){
	var s uint64 = 0
	var t uint64 = 1

	for i := 0; i < 64; i++{
		s += t
		t *= 2
	}
	f.Printf("Cabem %d grãos de milho em um tabuleiro de xadrez ", s)
}