programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto --> txt
  inclua biblioteca Tipos--> tp
funcao inicio() {

inteiro horas
escreva("")
leia (horas)

inteiro divisao_3h = horas / 3
inteiro resto = horas % 3

real valor = (divisao_3h * 10 ) + (resto * 5)   
real valor_final = mat.arredondar(valor, 2)

escreva ("O VALOR A PAGAR E = R$ ", valor_final, "\n")
  }
}
