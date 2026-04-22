// 36. Faça um programa que leia um número inteiro positivo na base 10, calcule e imprima o seu equivalente na base 16.
package main

import( 
	f "fmt"
	"strconv"
)

func main(){
	x := 0
	f.Print("Digite um número ")
	f.Scan(&x)

	y:= arrumar(hexadecimal(x))
	for _, z := range y {
		f.Printf("%s", z)
	}
}

func hexadecimal(num int) []string{
	var arr []string 
	r := 0

	for num > 0{
		r = num % 16
		arr = append(arr, traduzir(r))
		num /= 16
	}
	return arr
}

func traduzir(num int) string{
	if num <10{
		return strconv.Itoa(num)
	}
	if num == 10{
		return "A"
	}else if num == 11{
		return "B"
	}else if num == 12{
		return "C"
	}else if num == 13{
		return "D"
	}else if num == 14{
		return "E"
	}else{
		return "F"
	}

}

func arrumar(arr []string) []string{
	if len(arr) <= 1{
		return arr
	}

	primeiro := 0
	ultimo := len(arr) -1
	arr[primeiro], arr[ultimo] = arr[ultimo], arr[primeiro]

	arrumar(arr[1:ultimo])

	return arr
}