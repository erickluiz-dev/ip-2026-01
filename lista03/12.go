package main

import f "fmt"

func main(){
	X := 0.0
	f.Print("Digite um número ")
	f.Scan(&X)

	S := 0.0

	for i := 0; i < 20; i++{
		if (i % 2) == 0{
			S += X / fact(i)
		}else{
			S -= X / fact(i)
		}
	}
	f.Printf("A soma é: %.4f ", S)
}
func fact(num int) float64{
	if num == 0{
		return 1
	}
	return float64(num) * fact(num - 1)
}