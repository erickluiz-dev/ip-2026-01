package main

import(
	f "fmt"
	m "math"
)

func main(){
	a := 0.0
	f.Print("\n     ---Tabela---")
	for a<=6.3{	
		x := sen(a)
		f.Printf("Ângulo = %.1f | Seno = %f \n", a, x)
		a += 0.1
	}
}

func sen(a float64) float64{
	sen := a - (m.Pow(a, 3)/6) + (m.Pow(a, 5)/120) - (m.Pow(a, 7)/ 5040)
	return sen
}