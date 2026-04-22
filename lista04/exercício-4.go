package main 

import f "fmt"

func main(){
	decimal, j := 0, 0
	f.Print("Digite um número \n")
	f.Scan(&decimal)

	bin := binario(decimal)
	for i:= len(bin) - 1; i>=0; i--{
		f.Printf("%d ", bin[i])
	}
}

func binario(num int) []int{
	var arr []int
	x := 0
	for num >0{
		x = num % 2
		arr = append(arr, x)
		num /= 2
	}
	return arr
}