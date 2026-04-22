// 1. Escreva um programa que calcule potências. O usuário deve digitar a base e o expoente, e o programa deve
// apresentar o resultado (sem usar o comando pow). Assuma que o usuário irá digitar valores positivos.

package main 

import f "fmt"

func main(){
    base, expoente, resultado := 0.0, 0, 1.0
    f.Print("Digite a base ")
    f.Scan(&base)    
    f.Print("Digite o expoente ")
    f.Scan(&expoente)
    for i := 0; i < expoente; i++{
        resultado *= base
    }
    f.Printf("%.0f elevado à %d é %.0f ", base, expoente, resultado)
}
