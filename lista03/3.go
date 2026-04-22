// 3. Faça um programa que receba o salário de um funcionário chamado Carlos. Sabe-se que outro funcionário,
// João, tem salário equivalente a um terço do salário de Carlos. Carlos aplicará seu salário integralmente na
// caderneta de poupança, que está rendendo 2% ao mês, e João aplicará seu salário integralmente no fundo de
// renda fixa, que está rendendo 5% ao mês. O programa deverá calcular e mostrar a quantidade de meses
// necessários para que o valor pertencente a João iguale ou ultrapasse o valor pertencente a Carlos.
package main

import f "fmt"

func main() {
	salarioC := 0.0
	meses := 0
	f.Print("Digite o salário de carlos ")
	f.Scan(&salarioC)
	salarioJ := salarioC / 3

	for salarioC >= salarioJ {
		salarioC *= 1.02
		salarioJ *= 1.05
		meses++
	}
	f.Printf("OS valores se igualam em: %d meses ", meses)
}