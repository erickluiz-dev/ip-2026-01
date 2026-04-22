package main

import(
	f "fmt"
	m "math"
)

func main(){
	x:= 0.0
	f.Print("Digite um valor para X ")
	f.Scan(&x)

	diferença := cosseno(x) - m.Cos(x)
	y := cosseno(x)
	f.Printf("Cos(x) = %f | Diferença = %.2f ", y, diferença)
}

func cosseno(num float64) float64{
	cos := 1.0
	j := 2.0 
	for i:= 1; i < 20; i++{
		if i%2 == 0{
			cos += m.Pow(num, j) / fatorial(j)
		}else{
			cos -=m.Pow(num, j) / fatorial(j)
		}
		j += 2
	}  
	return cos
}

func fatorial(num float64) float64{
	if num <= 1{
		return 1
	}
	return num * fatorial(num-1)
}