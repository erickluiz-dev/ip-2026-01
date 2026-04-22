// 2. Elabore um programa que apresente os resultados da soma e da média aritmética dos valores pares situados na
// faixa numérica de 50 a 70.
package main 

import f "fmt"

func main(){
    soma, v, media := 0, 0, 0.0
    for i := 50; i <= 70; i++{
        if i % 2 == 0{
            soma += i
            v++
        }
    }
    media = float64(soma) / float64(v)
    f.Printf("A soma é: %d\n",soma)
    f.Printf("A média é: %.0f", media)
}