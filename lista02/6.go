package main

import f "fmt"

func main(){
	x, y := 0, 0
	f.Print("Digite o dividendo ")
	f.Scan(&x)
	f.Print("Digite o divisor ")
	f.Scan(&y)
	r := rest(x, y)
	if r == 0{
		f.Printf("O número %d é divisível por %d ", x, y)
	}else{
		f.Printf("O número %d não é divisível por %d ", x, y)	
	}
}
func rest (number1, number2 int) int {
	return number1 % number2
}