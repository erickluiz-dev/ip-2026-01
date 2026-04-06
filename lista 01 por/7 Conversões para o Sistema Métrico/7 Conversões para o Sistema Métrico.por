programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto --> txt
  inclua biblioteca Tipos --> tp
  funcao inicio() {
real temperatura_far
escreva ("")
leia (temperatura_far)

real polegadas
escreva ("")
leia (polegadas)

real temperatura_ce = (temperatura_far * 5 -160 ) / 9.0
real temperatura_final = mat.arredondar(temperatura_ce, 2)
cadeia temperatura = tp.real_para_cadeia(temperatura_final)
se (txt.posicao_texto(".", temperatura, 0) ==-1 ){
temperatura = temperatura + ".00"
}

real milimetros = polegadas * 25.4
real milimetros_final = mat.arredondar(milimetros, 2)
cadeia milimetros_f = tp. real_para_cadeia(milimetros_final)
se (txt.posicao_texto(".", milimetros_f, 0) ==-1){
milimetros_f = milimetros_f + ".00"
}

escreva ("O VALOR EM CELSIUS = ", temperatura, "\n")
escreva ("A QUANTIDADE DE CHUVA E = ", milimetros_f)
  }
}
