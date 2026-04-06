programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto--> txt
  inclua biblioteca Tipos --> tp
  funcao inicio() {

inteiro horas
escreva ("")
leia (horas)

inteiro minutos
escreva ("")
leia (minutos)

inteiro segundos
escreva ("")
leia (segundos)

inteiro resultado = (horas * 3600) + (minutos * 60) +segundos

escreva ("O TEMPO EM SEGUNDOS E = ", resultado, "\n")
  }
}
