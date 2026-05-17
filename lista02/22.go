// 22. Desenvolver um programa que calcule o salário bruto e o salário líquido de um funcionário.
// • Dados de Entrada: Matrícula do funcionário (int);
// Quantidade de horas-extras trabalhadas.

// • Constantes: Salário Mínimo = R$ 788.00;
// Valor da Hora-Extra = R$ 10.00.

// Sabe-se:
// • Salário hora-extra = horas-extras * Valor da Hora-Extra;
// • Salário bruto = 3 * Salário Mínimo + Salário hora-extra;
// • Desconto INSS = 12 % do salário bruto, se salário bruto for maior que R$ 1500,00;
// • Desconto do Imposto de Renda = 20 % do Salário Bruto se o mesmo for maior que R$ 2000,00;

package main

import f "fmt"


func main(){
	matricula := 0
	f.Print("Digite matrícula do funcionário ")
	f.Scan(&matricula)
	h_extras := 0
	f.Print("Digite a quantidade de horas extras trabalhadas ")
	f.Scan(&h_extras)
	s_minimo := 788
	sb := salario_bruto(h_extras, s_minimo)

	sl := sb
	if sl > 1500 && sl > 2000{
		sl *= 0.67 
	}else if sl > 1500{ 
		sl *=  0.88
	}
	f.Printf("Matrícula: %d \nSalário bruto: R$ %.2f \nSalário líquido: R$ %.2f", matricula, sb, sl)
}
func salario_bruto(numero1, numero2 int) float64{
	s := 10 * numero1 
	s += 3 * numero2
	return float64(s)
}