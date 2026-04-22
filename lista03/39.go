// 39. Num frigorífico existem 90 bois. Cada boi traz preso no seu pescoço um cartão contendo um número de identificação e seu peso. Implementar 
// um programa que escreva o número e o peso do boi mais gordo e do boi mais magro (não é necessário armazenar os dados de todos os bois).
package main

import f "fmt"

func main(){
	maior, menor := 0.0, 0.0
	IdMenor, idMaior := 0.0, 0.0

	for i:=0; i<90; i++{

		identificação, peso :=0.0, 0.0
		f.Print("Identificação: ")
		f.Scan(&identificação)
		f.Print("Peso: ")
		f.Scan(&peso)

		if i==0{
			menor = peso
			IdMenor = identificação
		}
		if peso > maior{
			maior = peso
			idMaior = identificação
		}
		if peso < menor{
			menor = peso
			IdMenor = identificação
		}
	}
	f.Printf("Maior boi:\n Identificação: %.0f Peso: %.2f \n Menor boi: \n Identificação: %.0f Peso: %.2f", idMaior, maior, IdMenor, menor)

}