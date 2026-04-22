package main

import f "fmt"

func main(){
	x, y := 100.0, 0.0
	res := 0.0

	for i := 0; i< 20; i++{
		y = float64(i)
		res += x / factorial(y)
		x--
	}
	f.Printf("A soma é: %f ", res)
}
func factorial(n float64) float64{
	if n == 0{
		return 1
	}
	return n * factorial(n-1)
}