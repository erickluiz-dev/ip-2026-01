package main

import f "fmt"

func main(){
	cpf := 0
	f.Print("Digite seu CPF ")
	f.Scan(&cpf)

	NumVer := calculo(Decompor(cpf)) 
	for _, v := range NumVer {
		f.Printf("%.f ", v)
	}
}

func Decompor(num int) []float64{
	var arr []float64

	for num > 0{
		arr = append(arr, float64(num % 10))
		num /= 10
	}
	return trocar(arr)
}

func calculo(arr []float64) []float64{

	s := 0.0
	a := 0

	for i:= (len(arr) + 1); i>=2; i--{
		s += arr[a] * float64(i)
		a++
	}
	s = resto(s)
	
	if s < 2{
		s = 0
	}else{
		s = 11 - s
	}
	arr = append(arr, s)
	if len(arr) == 11{
		return arr
	}
	return calculo(arr)


}
func resto(num float64) float64{
	res := num
	for res >= 11{
		res -= 11
	}
	return res
}
func trocar(arr []float64) []float64 {
	for i, j := 0, len(arr)-1; i < j; i, j = i+1, j-1 {
		arr[i], arr[j] = arr[j], arr[i]
	}
	return arr
}