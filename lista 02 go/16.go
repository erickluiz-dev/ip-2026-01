// 16. Desenvolver um programa que leia os coeficientes (A, B e C) de uma equação do segundo grau ( Ax2 + Bx + C
// =0) e que calcule suas raízes. O programa deve mostrar, quando possível, o valor das raízes calculadas e a 
// classificação das mesmas: “RAÍZES IMAGINÁRIAS”, “RAIZ ÚNICA” ou “RAÍZES DISTINTAS”
package main

import(
	 f "fmt"
	 "math"
)

func main(){
	a, b, c := 0.0, 0.0, 0.0
	f.Println("Digite os coeficientes A B C ")
	f.Scan(&a, &b, &c)
	r1, r2 := equacao(a, b, c)
	if math.IsNaN(r1) && math.IsNaN(r2){
	 f.Printf("Raízes imaginárias ")
	}else if r1 == r2{
		f.Printf("Raís única S = {%.2f} ", r1)
	}else{
		f.Printf("Raízes distintas S = {%.2f, %.2f} ", r1, r2) 
	}
}
func equacao(A, B, C float64) (float64, float64) { 

	delta := (B * B) - 4 * A * C
	raizdelta := math.Sqrt(delta)
	bhaskara1 := (-B + raizdelta) / (2 * A)
	bhaskara2 := (-B - raizdelta) / (2 * A)
	return bhaskara1, bhaskara2
}