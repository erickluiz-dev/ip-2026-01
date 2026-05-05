package main

import f "fmt"

func main(){
    var arr [10]int
    contagem := make(map[int]int)
    for i := 0; i < len(arr); i++{
        f.Scan(&arr[i])
        contagem[arr[i]]++
    }
    f.Print("---Números Repetidos---\n")

    for i, v := range contagem{
        if v > 1{
            f.Printf("Número: %d | Repetições: %d \n", v, i)
        }
    }
}