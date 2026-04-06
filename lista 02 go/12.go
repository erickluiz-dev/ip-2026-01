// Idade Classificação
// 0 a 2 anos Recém-nascido
// 3 a 11 anos Criança
// 12 a 19 anos Adolescente
// 20 a 55 anos Adulto
// Acima de 55 anos Idoso
package main 

import f "fmt"

func main(){
	x := 0
	f.Print("Digite sua idade ")
	f.Scan(&x)
	a := classification(x)
	f.Printf("A idade é %d a classificação é %s ", x, a )
}
func classification(number int) string {
	age := "age"
	if number >= 0 && number <=2 {
		age = "Recém-nascido" 
	}else if number >= 3 && number <= 11 {
		age = "Criança" 
	}else if number >= 12 && number <= 19 {
		age = "Adolescente" 
	}else if number >= 20 && number <= 55 {
		age = "Adulto" 
	}else if number > 55{
		age = "Idoso" 
	}
	return age
}