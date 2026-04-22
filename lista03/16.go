package main

import f "fmt"

func main(){
	a1, a2, n := 0, 0, 0
	f.Print("Digite um valor para A1 ")
	f.Scan(&a1)
	f.Print("Digite um valor para A2 ")
	f.Scan(&a2)
	f.Print("Digite o número de termos ")
	f.Scan(&n)

	f.Println()
	f.Printf("%d - %d - ", a1, a2)
	ai := 0
	for i:=3; i<=n; i++{
		if i%2 == 0{
			ai = a1 - a2
		}else{
			ai = a1 + a2	
		}
		f.Printf("%d - ", ai)
		a1 = a2
		a2 = ai
	}
}