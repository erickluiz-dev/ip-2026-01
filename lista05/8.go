// Escreva um programa que receba quinze números inteiros e armazene em um vetor a raiz quadrada de cada 
// número. Caso o valor digitado seja menor do que zero, o número -1 deve ser atribuído ao elemento do vetor.
// Após isso, imprima todos os valores armazenados.

package main

import(
	f "fmt"
	m "math"
)

func main(){
	arr := make([]float64, 15)
	for i := 0; i < len(arr); i++{
		y:= 0
		f.Scan(&y)
		arr[i] = roots(y)
	}
	f.Printf("Resultados: %v", arr)
}
func roots(x int) float64{
	if x < 0 {
		return -1	
	}else{
		return m.Sqrt(float64(x))
	}
}
