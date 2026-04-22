// 35. Faça um programa que leia um número inteiro positivo na base 10, calcule e imprima o seu equivalente na
// base 2.
package main

import f "fmt"

func main(){
	x := 0
	f.Print("Digite um número ")
	f.Scan(&x)
	
	y:= organizar(binario(x))
	f.Printf("%d ", y)
		
}
func binario(num1 int) []int{
	var bin []int
	r:= 0
	for num1 > 0{
		r = num1 % 2 
		bin = append(bin, r)
		num1 /= 2
	}
	return bin
}

func organizar(slice []int ) []int{
	if len(slice) <= 1{
		return slice
	}
	primeiro := 0
	ultimo := len(slice) - 1
	slice[primeiro], slice[ultimo] = slice[ultimo], slice[primeiro]
	
	organizar(slice[1:ultimo])

	return slice
}