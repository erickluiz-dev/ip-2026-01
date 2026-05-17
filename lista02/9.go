package main

import f "fmt"

func main(){	
	x := 0.0
	f.Print("Digite o valor da compra ")
	f.Scan(&x)
	s, e := sale(x)
	if !e {
		f.Printf("O valor da venda é %.2f ", s)
	}else{
		f.Printf("Compra invalida ")
	}
}
func sale (number float64) (float64, bool){
	erro := false 
	sale := 1.0
	if number <= 0{
		erro = true
	}
	if number < 10{
		sale = number * 1.7
	}else if number >= 10 && number < 30{
		sale = number * 1.5
	}else if number >= 30 && number < 50{
		sale = number * 1.4
	}else{
		sale = number * 1.3
	}
	return sale, erro
}