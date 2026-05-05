package main

import(
	f "fmt"
	"math/rand"
)

func main(){
	Map := make(map[int]int)
	for i := 0; i < 20; i++{
		Map[rand.Intn(6)+1]++
	}
	f.Println("---Números e Frequência---")

	for i, v := range Map{
		f.Printf("Número: %d | Repetições: %d \n", i, v)
	}
}