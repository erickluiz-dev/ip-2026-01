// 30. Faça um programa que calcule o volume de uma esfera em função do raio R. O raio deverá variar de 0 a 20 cm de 0,5 em 0,5 cm.
package main

import(
	 f "fmt"
	 m "math"
)

func main(){
	f.Println()
	f.Print("     ---Tabela---\n")

	for i := 0.0; i <= 20.0; i += 0.5{
		v := volume(float64(i))
		f.Printf("Volume = %.f cm³ | Raio = %.1f\n", v, i)
	}	
}
func volume(num float64) float64{
	return (4.0/3.0) * m.Pi * m.Pow(num, 3)
}