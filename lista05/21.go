package main

import f "fmt"

func main(){

	type conjunto struct{
		codigo int
		array []float64
	}

	Conjunto := []conjunto{
		{codigo: 0, array: []float64{}},
	}

	for{
		Conjunto[0].array = []float64{}

		cod := 0
		f.Print("Digite o código: ")
		f.Scan(&cod)
		Conjunto[0].codigo = cod
		
		if Conjunto[0].codigo == 0{
			break
		}

		f.Println("Digite 10 números reais: ")

		for i := 0; i < 10; i++{
			x := 0.0
			f.Scan(&x)
			Conjunto[0].array = append(Conjunto[0].array, x)
		}

		if Conjunto[0].codigo == 1{
			f.Print(Conjunto[0].array)
		}

		if Conjunto[0].codigo == 2{
			f.Print(inverter(Conjunto[0].array[:]))
		}

	}
}	

func inverter(arr []float64) []float64{
	if len(arr) <= 1{
		return arr
	}
	primeiro := 0
	ultimo := len(arr)-1

	arr[primeiro], arr[ultimo] = arr[ultimo], arr[primeiro]

	inverter(arr[1:ultimo])

	return arr
}