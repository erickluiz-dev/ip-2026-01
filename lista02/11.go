package main

import f "fmt"

func main(){
	x := 0.0
	f.Print("Digite o x da funação ")
	f.Scan(&x)
	y := funcao(x)
	f.Printf("O y da função é %.2f ", y)
}
func funcao(x float64) float64 {
	y := 8 / (2-x)
	return y
}