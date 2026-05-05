// 12. Faça um programa que: - leia um conjunto de valores inteiros correspondentes a 15 notas dos alunos de uma turma. 
// As notas variam de 0 a 10; - calcule a frequência absoluta e a frequência relativa de cada nota; - imprima uma tabela contendo 
// os valores das notas (0 a 10) e suas respectivas frequências absoluta e relativa. 
package main

import f "fmt"

func main(){
	notas := make(map[int]int)
	var arr [5]int
	for i := 0; i < len(arr); i++{
		f.Scan(&arr[i])
		notas[arr[i]]++
	}
	for i := 0; i <= 10; i++{
		f.Printf("Nota: %d | Frequência absoluta: %d | Frequência relativa: %.1f \n", i, notas[i], (float64(notas[i]) / float64(len(arr))))
	}
}