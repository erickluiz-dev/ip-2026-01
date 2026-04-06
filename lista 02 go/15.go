// 15. Faça um programa que leia uma data (dia, mês e ano, em uma variável inteira cada), e escreva a mesma data
// no formato dia de (mês por extenso) de ano.
package main 

import f "fmt"

func  main(){
	day, month, year :=0, 0, 0
	f.Println("Digite uma data ")
	f.Scan(&day, &month, &year)
	var strmonth string
	if month ==1{
		strmonth = "janeiro"
	}
	if month ==2{
		strmonth = "fevereiro"
	}
	if month ==3{
		strmonth = "março"
	}
	if month ==4{
		strmonth = "abril"
	}
	if month ==5{
		strmonth = "maio"
	}
	if month ==6{
		strmonth = "junho"
	}
	if month ==7{
		strmonth = "julho"
	}
	if month ==8{
		strmonth = "agosto"
	}
	if month ==9{
		strmonth = "setembro"
	}
	if month ==10{
		strmonth = "outubro"
	}
	if month ==11{
		strmonth = "novembro"
	}
	if month ==12{
		strmonth = "dezembro"
	}
	f.Printf("A data é %d de %s de %d ", day, strmonth, year)
}