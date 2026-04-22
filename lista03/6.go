// 6. Escreva um programa que receba um número inteiro positivo, verifique e informe se ele é ou não um número
// triangular. Obs.: Um número é triangular quando é resultado do produto de três números naturais consecutivos.
// Exemplo: 24 = 2 x 3 x 4; 120 = 4 x 5 x 6
package main

import f "fmt"


func main() {
    number := 0
    f.Print("Digite um número ")
    f.Scan(&number)


    verification := true
    x, y, z := 1, 2, 3


    for z <= number / 2 {
        if (x * y * z) == number {
            f.Printf("%d é um número triângular \n", number)
            verification = false
            break
        }
        x++
        y++
        z++
    }
    if verification {
        f.Printf("%d não é um número triângular \n", number)
    }
}