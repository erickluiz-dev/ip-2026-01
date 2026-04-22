// 18. Escreva um programa que imprima os índices da diagonal principal de uma matriz 10x10.
package main

import (
	f "fmt"
)

func main(){
	for i:=0; i<10; i++{
		for j:=0; j<10; j++{
			if i == j{
				f.Printf("[%d][%d] ", i, j)
			}else{
				f.Print("       ")
			}
		}
		f.Println("")
	}
}