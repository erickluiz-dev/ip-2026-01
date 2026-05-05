package main

import f "fmt"

func main(){
	conta := make(map[int]float64)

	for i := 0; i < 10; i++{
		var codigo int
		f.Print("Digite o número da conta: ")
		f.Scan(&codigo)

		if conta[codigo] > 0{
		f.Println("Número da conta inválido")
		f.Print("Digite o número da conta: ")
		f.Scan(&codigo)
		} 

		var saldo float64
		f.Print("Digite o saldo da conta: ")
		f.Scan(&saldo)

		conta[codigo] = saldo
	}

	finalizar := false
	for{
		if !finalizar{
			conta, finalizar = menu_principal(conta) 
		}else{
			break
		}
	}

}
func menu_principal(conta map[int]float64) (map[int]float64, bool){

	f.Println("\n---Menu Principal---")

	var opcao int
	f.Print("1. Efetuar depósito\n2. Efetuar saque\n3. Consultar o ativo bancário (ou seja, o somatório dos saldos de todos os clientes)\n4. Finalizar o programa.\n")
	f.Scan(&opcao)

	switch opcao{
		case 1:
			var codigo int	
			f.Println("Digite o código da conta: ")
			f.Scan(&codigo)

			var valor float64
			f.Println("Digite o valor a ser depositado: ")
			f.Scan(&valor)

			if conta[codigo] < 1{
				f.Print("Conta não encontrada \n")
				return conta, false
			}else{
				conta[codigo] += valor
				return conta, false
			}

			case 2:
			var codigo int	
			f.Println("Digite o código da conta: ")
			f.Scan(&codigo)

			var valor float64
			f.Println("Digite o valor a ser sacado: ")
			f.Scan(&valor)

			if conta[codigo] < 1{
				f.Print("Conta não encontrada \n")
				return conta, false
			}else if (conta[codigo]- valor) < 0{
				f.Print("Saldo insuficiente \n")
				return conta, false
			}else{
				conta[codigo] -= valor
				return conta, false
			}

			case 3:
				var soma float64 
				for _, v := range conta{
					soma += v
				}
				f.Printf("Ativo bancário: R$ %.2f \n", soma) 
				return conta, false
			
			case 4: 
			return conta, true

			default: 
			f.Print("Opção inválida ")
			return conta, false
	}


}