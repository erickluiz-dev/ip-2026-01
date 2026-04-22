// 21. Escreva um programa que calcule e imprima o valor de bn. O usuário vai informar os valores de b e n. Assuma
// que o valor de n é maior do que 1 e o valor de b é maior ou igual a 2, ambos valores
package main

import f "fmt"

func main() {
    b, n := 0, 0
    f.Print("Digite o valor de B e N \n")
    f.Scan(&b)
    f.Scan(&n)

    r := Pow(b, n)
    f.Printf("%d ", r)
}
func Pow(n1, n2 int) int {
    res := 1
    for i := 0; i < n2; i++ {
        res *= n1
    }
    return res
}
