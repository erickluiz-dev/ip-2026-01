package main 

import f "fmt"

func main(){
	H := 0.0

	for i := 1.0; i <= 50; i++{
		H += (i*2 - 1) / i 
	}
	f.Printf("H = %f ", H)
}