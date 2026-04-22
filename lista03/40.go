// 40. Uma Companhia de teatro planeja dar uma série de espetáculos. A direção calcula que, a R$ 6,00 o ingresso, serão vendidos 130 ingressos e as 
// despesas montarão em R$300,00. A uma diminuição de R$ 0,60 no preço dos ingressos espera-se que haja um aumento de 30 ingressos vendidos. Fazer
// um programa que escreva uma tabela de valores do lucro esperado em função do preço do ingresso, fazendo-se variar este preço de R$ 6,00 a R$ 1,00
//  de R$ 0,60 em R$ 0,60. Escreva ainda o lucro máximo esperado, o preço e o número de ingressos correspondentes.
package main 

import f "fmt"

func main(){
	preço := 6.0
	entrada:=0.0
	lucro := 0.0
	LucroMax, PreçoMax, NumMax := 0.0, 0.0, 0

	f.Print("\n                ---Tabela---\n")
	for i:=0; preço > 1.0; i++{
		entrada = preço * (130 + (30 * float64(i)))
		lucro = entrada-300  
		f.Printf("Preço do Ingresso: R$ %.2f --> Lucro: R$ %.2f \n", preço, lucro)
		if lucro > LucroMax{
			LucroMax = lucro 
			PreçoMax = preço
			NumMax = 130 + (30 * i)
		}
		preço -= 0.6
	}
	f.Println()
	f.Printf("Lucro máximo: R$ %.2f | Preço: R$ %.2f | Número de Ingressos: %d", LucroMax, PreçoMax, NumMax)
	f.Println("\n")
}	