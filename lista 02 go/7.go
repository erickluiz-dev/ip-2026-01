package main 

import f "fmt"

func main(){
	a, b, c := 0, 0, 0
	f.Print("Digite o 1º valor ")
	f.Scan(&a)
	f.Print("Digite o 2º valor ")
	f.Scan(&b)
	f.Print("Digite o 3º valor ")
	f.Scan(&c)

	menor, inter, maior := organization(a, b, c)
	f.Printf("Esses números em orde crescente são %d, %d e %d ", menor, inter, maior)
}
func organization(number1, number2, number3 int) (int, int, int) {
	aux := 0
	if number1 > number2{
		aux = number2
		number2 = number1
		number1 = aux
	}
	if number1 > number3{
		aux = number3
		number3 = number1
		number1 = aux
	}
	if number2 > number3{
		aux = number3
		number3 = number2
		number2 = aux
	}
	return number1, number2, number3
}