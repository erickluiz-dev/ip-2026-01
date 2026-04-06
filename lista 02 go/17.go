// 17. Desenvolver um programa para calcular a conta de água para a SANEAGO. O custo da água varia dependendo
// do tipo do consumidor - residencial, comercial ou industrial. A regra para calcular a conta é:
// • Residencial: R$ 5,00 de taxa mais R$ 0,05 por m3 gastos;
// • Comercial: R$ 500,00 para os primeiros 80 m3 gastos mais R$ 0,25 por m3 gastos acima dos 80 m3;
// • Industrial: R$ 800,00 para os primeiros 100 m3 gastos mas R$ 0,04 por m3 gastos acima dos 100 m3;
// O programa deverá ler a conta do cliente, seu tipo (residencial, comercial e industrial) e o seu consumo de água em 
// metros cúbicos. Como resultado imprimir a conta do cliente e o valor em real a ser pago pelo mesmo
package main

import f "fmt"
func main(){
	var account string
	var type_ string
	spent := 0.0
	f.Println("Digite sua conta o tipo de consumidor e seu gasto de m³ ")
	f.Scan(&account, &type_, &spent)
	e, t := end_spent(type_, spent)
	if !t {
	f.Printf("Conta %s valor a ser pago = R$ %.2f ", account, e)
	}else{
		f.Print("Tipo de conta invalido ")
	}
}
func end_spent(word string, number float64) (float64, bool){
	start_spent := 0.0
	test := false
	if word == "residencial"{
		start_spent = 5 
		start_spent += 0.05 * number
	}else if word == "comercial"{
		start_spent = ((number - 80) * 0.25) + 500
	}else if word == "industrial"{
		start_spent = ((number - 100) * 0.04) + 800
	}else{
		test = true
	}
	return start_spent, test
}