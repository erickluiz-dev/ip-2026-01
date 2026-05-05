package main

import f "fmt"

func main(){
	var Num [10]int
	f.Println("---Vetor Número---")
	for i := 0; i < len(Num); i++{
		f.Scan(&Num[i])
	}

	var Divis [5]int
	f.Println("---Vetor Divisor---")
	for i := 0; i < len(&Divis); i++{
		f.Scan(&Divis[i])
	}

	for _, v := range Num{
		f.Printf("Número %d: \n", v)
		for i, u := range Divis{
			if v % u == 0{
				f.Printf("Divisível por %d na posição %d\n", v, i)
			}
		}
	}
}