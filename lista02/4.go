package main 

import(
	 f "fmt"
	 "math"
)

func main(){
	x, r := 0.0, 0.0
	f.Print("Digite um valor ")
	f.Scan(&x)
	if x >= 0{
		r = squareroot(x)
		f.Printf("A raiz quadrada de %.2f é %.2f ", x, r)
	}else{
		r = pow(x)
		f.Printf("O quadrado de %.2f é %.2f ", x, r)
	}
}
func squareroot(number float64) float64 {
	return math.Sqrt(number)
}
func pow(number float64) float64 {
	return math.Pow(number, 2 )
}