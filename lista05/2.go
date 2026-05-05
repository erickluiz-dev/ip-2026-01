package main 

import f "fmt"

func main(){
    primeiro := make([]int, 10)
    var vetorpar []int
    var vetorimpar []int
    
    for i := 0; i < 10; i++{
        f.Scan(&primeiro[i])
        if primeiro[i] % 2 == 0{
            vetorpar = append(vetorpar, primeiro[i])
        }else{
            vetorimpar = append(vetorimpar, primeiro[i])
        }
    }
    
    soma:= 0
    segundo := make([]int, 5)
    for i := 0; i < 5; i++{
        f.Scan(&segundo[i])
        soma += segundo[i]
    }
   
   var primeiro_vetor_resultante []int
   for i := 0; i < len(vetorpar); i++{
       calculo := vetorpar[i] + soma
       primeiro_vetor_resultante = append(primeiro_vetor_resultante, calculo)
   }
   var segundo_vetor_resultante []int
   for i := 0; i < len(vetorimpar); i++{
       calculo := vetorimpar[i]+soma
       segundo_vetor_resultante = append(segundo_vetor_resultante, calculo)
   }
   for _, v := range primeiro_vetor_resultante{
       f.Printf("%d ", v)
   }
   f.Println()
    for _, v := range segundo_vetor_resultante{
       f.Printf("%d ", v)
   }
}
