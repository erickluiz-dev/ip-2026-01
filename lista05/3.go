
package main

import f "fmt"

func main(){
    var arr [10]int
    for i := 0; i < len(arr); i++{
        f.Scan(&arr[i])
    }
    f.Print("\n---Números Pares---\n")
    for _, v := range arr{
        if v % 2 == 0{
            f.Printf("%d ", v)
        }
    }
    
        f.Print("\n---Soma dos Números Pares---\n")
        soma_pares := 0
    for _, v := range arr{
        if v % 2 == 0{
            soma_pares += v
        }
    }
    f.Printf("%d ", soma_pares)
    
    f.Print("\n---Números Ímpares---\n")
    for _, v := range arr{
        if v % 2 != 0{
            f.Printf("%d ", v)
        }
    }
    
    f.Print("\n---Soma dos Números Ímpares---\n")
    soma_impares := 0
    for _, v := range arr{
        if v % 2 != 0{
            soma_impares += v
        }
    }
    f.Printf("%d ", soma_impares)
