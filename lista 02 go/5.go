package main

import f "fmt"

func main(){
	x := 0
	f.Println("Digite um número ")
	f.Scan(&x)
	r := x % 5
	if r == 0{
		f.Printf("O número %d é divisível por 5 ", x)
	}else{
		f.Printf("O número %d não é divisível por 5 ", x)
	}
}