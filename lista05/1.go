// Faça um programa que preencha um vetor com 10 números inteiros. Calcule e mostre os números superiores a 50 e suas respectivas
// posições. O programa deverá mostrar uma mensagem se não existir nenhum número nessa condição.


package main

import f "fmt"

func main() {
    var arr [10]int
    for i := 0; i < 10; i++ {
        f.Scan(&arr[i])
    }

    z := 0
    y := 0

    for i, v := range arr {
        if v > 50 {
            if z == 0 {
                f.Print("---Números superiores a 50---\n")
            }
            f.Printf("Número: %d | Posição: %d \n", v, i+1)
            z++
            y++
        }
        if y == 0 {
            f.Print("Não existe número maior que 50 \n")
            return
        }
    }
}
