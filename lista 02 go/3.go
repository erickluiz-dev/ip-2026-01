package main

import f "fmt"

func main(){
	
	x, y, result := 0, 0, 0
	f.Print("Digite o primeiro número ")
	f.Scan(&x)
	f.Print("Digite o segundo número ")
	f.Scan(&y)
	s := soma(x, y)
	if s > 20{
		result = s + 8
	}else{
		result = s - 5
	}
	f.Printf("O resulta da soma entre %d e %d é %d ", x, y, result)
}
func soma(number1, number2 int) int{
	return number1 + number2  	
}