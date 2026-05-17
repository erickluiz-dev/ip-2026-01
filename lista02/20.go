//20. Elabore um programa que calcule o valor a ser pago por um produto considerando o preço normal de etiqueta
//e a escolha da condição de pagamento. Utilize os códigos da tabela a seguir para saber qual a condição de
//pagamento escolhida e efetuar o cálculo adequado.
//Código Condição de Pagamento
//1 À vista, dinheiro ou cheque, 10% de desconto
//2 À vista, cartão de crédito, 5% de desconto
//3 Em 2 vezes, preço normal de etiqueta sem juros
//4 Em 3 vezes, preço normal de etiqueta + 10% de juros

package main

import f "fmt"

func main(){
	buy := 0.0 
	f.Print("Qual é o valor da compra? ")
	f.Scan(&buy)
	
	option := 0
	f.Print("Escolha uma opção de pagamento \n1. À vista, dinheiro ou cheque \n2. À vista, cartão de crédito \n3. Em 2 vezes \n4. Em 3 vezes \n")
	f.Scan(&option)
	if option < 1 || option > 4{
		f.Print("Opção de pagamento invalida ")
		return
	}
	
	if option ==1{
		buy *= 0.9
	}else if option ==2{
		buy *= 0.95
	}else if option ==4{
		buy *=  1.1
	}
	f.Printf("O valor a ser pago pelo produto é R$ %.2f ", buy)
}