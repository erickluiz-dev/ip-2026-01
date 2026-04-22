package main

import(
	f "fmt"
	m "math"
)

func main(){
	x := soma(0)
	f.Printf("∏ = %f ", x)
}

func soma(num int) float64{
	S := float64(num)
	j := 1.0                     //para a divisão
	for i := 1; i <= 51; i++{    //para saber se soma ou subtrai
		if i%2 == 0{
			S -= (1 / m.Pow(j, 3))
		}else{
			S += (1 / m.Pow(j, 3))
		}
		j += 2.0
	}
	return Calculo(S)
}

func Calculo(num float64) float64{
	y := m.Cbrt(num*32)
	return y
}