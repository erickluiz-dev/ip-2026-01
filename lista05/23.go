package main


import (
    f "fmt"
    "math/rand"
)


func main() {

    sorteio, aceitar := false, 0
    f.Println("Se dejesa o ônius com lugares prévios digite 1, senão digite 0 ")
    f.Scan(&aceitar)

    if aceitar == 1 {
        sorteio = true
    }

    var janela [24]int

    var corredor [24]int

    if sorteio {
        for i := 0; i < 24; i++ {
            num := rand.Intn(2)
            janela[i] = num
        }

        for i := 0; i < 24; i++ {
            num := rand.Intn(2)
            corredor[i] = num
        }
    }

    JanelaCheia, CorredorCheio := false, false

MenuVendas:
    for {
        f.Print("\n---Menu de Vendas---")

        somajanela := 0 // verifica se a janela está cheia
        for _, v := range janela {
            somajanela += v
        }
        if somajanela == 24 {
            JanelaCheia = true
        }

        somacorredor := 0 // verifica se o corredor está cheio
        for _, v := range corredor {
            somacorredor += v
        }
        if somacorredor == 24 {
            CorredorCheio = true
        }

        opcao := 0
        f.Print("\n1. Escolher poltrona no corredor\n2. Escolher poltrona na janela\n3. Encerrar o progama\n")
        f.Scan(&opcao)

        switch opcao {
        case 1:
            if JanelaCheia && CorredorCheio {
                f.Print("Sem lugares diponíveis ")
                continue MenuVendas
            } else if JanelaCheia {
                f.Print("Sem espaço na janela ")
                continue MenuVendas
            } else {
                f.Println("Poltronas disponíveis: ")


                for i, v := range janela {
                    if v == 0 {
                        f.Printf("Poltrona:%d | ", i+1)
                    }
                }

                lugar := 0
                f.Print("\nDigite a poltrona: ")
                f.Scan(&lugar)
                janela[lugar-1] = 1
            }

        case 2:
            if JanelaCheia && CorredorCheio {
                f.Print("Sem lugares diponíveis ")
                continue MenuVendas
            } else if CorredorCheio {
                f.Print("Sem espaço no corredor ")
                continue MenuVendas
            } else {
                f.Println("Poltronas disponíveis: ")

                for i, v := range corredor {
                    if v == 0 {
                        f.Printf("Poltrona:%d | ", i+1)
                    }
                }

                lugar := 0
                f.Print("\nDigite a poltrona: ")
                f.Scan(&lugar)
                corredor[lugar-1] = 1
            }

        case 3:
            break MenuVendas

        default:
            f.Print("Opção Inválida ")

        }
    }
}
