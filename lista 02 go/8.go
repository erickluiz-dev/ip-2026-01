package main

import f "fmt"

func main(){
	x := 0
	f.Print("Digite um valor ")
	f.Scan(&x)
	if x >= 20 && x <= 90 {
		f.Printf("O número %d está compreendido entre 20 e 90", x)
	}else{
		f.Printf("O número %d não está compreendido entre 20 e 90", x)
	}
}