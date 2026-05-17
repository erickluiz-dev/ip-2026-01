package main

import f "fmt"

func main(){
	x :=  0
	f.Print("Digite um número ")
	f.Scan(&x)
	if x < 0{
		f.Printf("O número %d é negativo ", x)
	}else if x > 0{
		f.Printf("O número %d é positivo ", x)
	}else{
		f.Printf("O número %d é nulo ", x)
	}  
}