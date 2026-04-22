// 7. Escreva um programa que receba vários números, calcule e mostre:
// a) a soma dos números digitados;
// b) a quantidade de números digitados;
// c) a média dos números digitados;
// d) o maior número digitado;
// e) o menor número digitado;
// f) a média dos números pares;
// g) a percentagem dos números ímpares entre todos os números digitados.
// Finalize a entrada de dados com a digitação do número 30.000.
package main 

import f "fmt"

func main(){
	numero := 0
	soma := 0
	quantidade := 0
	numeros_pares := 0
	soma_numeros_pares := 0
	maior_numero := 0
	menor_numero := 0
	qtd_impar := 0
	i := 0

	for{
		f.Print("Escreva um número ")
		f.Scan(&numero)
		if numero == 30000{
			break
		}
		if i == 0{
			maior_numero = numero
			menor_numero = numero
		}
		soma += numero
		quantidade++
		if (numero % 2) == 0{
			numeros_pares++
			soma_numeros_pares += numero
		}else{
			qtd_impar++
		}
		if numero > maior_numero{
			maior_numero = numero
		}
		if numero < menor_numero{
			menor_numero = numero
		}
		i++
	}
	media := float64(soma) / float64(quantidade)
	media_pares := float64(soma_numeros_pares) / float64(numeros_pares)
	porcentagem := float64(100 * qtd_impar) / float64(quantidade)

	f.Printf("Soma: %d\n", soma)
	f.Printf("Quantidade de números: %d\n", quantidade)
	f.Printf("Média dos números digitados: %.2f\n", media)
	f.Printf("Maior númmero: %d\n", maior_numero)
	f.Printf("Menor numero: %d\n", menor_numero)
	f.Printf("Média dos números pares: %.2f\n", media_pares)
	f.Printf("Porcentagem dos números ímpares: %.2f\n", porcentagem)
}