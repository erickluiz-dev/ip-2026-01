package main

import f "fmt"

func main(){
	var valor int
	f.Print("Digite um valor ")
	f.Scan(&valor)
	test := valor % 2
	if test == 0{
		f.Printf("O número %d é par ", valor)
	}else{
		f.Printf("O número %d é impar", valor)
	}
}