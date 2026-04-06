package main

import (
	f "fmt"
	m "math"
)

 func main(){
	figure := 0 // Definição da figuara
	f.Println("Escolha uma opção \n1-cone \n2-cilindro \n3-esfera")
	f.Scan(&figure)
	if figure != 1 && figure != 2 && figure != 3{
		f.Print("Opição invalida ")
		return	
	}
	cone, cylinder := false, false //verifica a escolha do usuário 
	if figure == 1{
		cone = true
	}
	if figure ==2{
		cylinder = true
	}

	radiu := 0.0 //declara o raio
	f.Print("Qual é o raio da figura? ")
	f.Scan(&radiu)

	height := 0.0 //declara a altura
	if cone || cylinder	{
		f.Print("Qual é a altura da figura? ")
		f.Scan(&height)
	}

	a, v := exit(cone, cylinder, radiu, height) // faz os calculos e escreve a área e o volume
	f.Printf("A área é = %.2f o volume é = %.2f ", a, v)
}
func exit(type1, type2 bool, number1, number2 float64) (float64, float64){
	r, h := number1, number2
	area, volume:= 0.0,  0.0
	if type1{                                 // faz os cálculos cone reto
		area = m.Pi * r * m.Sqrt((r * r)+ (h * h)) + (m.Pi * (r * r))
		volume = (m.Pi * (r * r) * h)/ 3
	}else if type2{                           // faz os cálculos do cilindro
		area = (2 * m.Pi * r * h) + (2 * m.Pi *(r * r))
		volume = (m.Pi * (r * r)) * h
	}else{                                   // faz os calculos da esfera
		area = 4 * m.Pi * (r * r)
		volume = (4  * m.Pi * (r * r * r)) /3
	}
	return area, volume
}