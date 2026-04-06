programa {
inclua biblioteca Matematica --> mat
inclua biblioteca Texto --> txt
inclua biblioteca Tipos --> tp
  funcao inicio() {

real a    
escreva ("")
leia (a)

real b 
escreva ("")
leia (b)

real c
escreva ("")
leia (c)

real d
escreva ("")
leia (d)

real det = (a * d) - (b * c)
det = mat.arredondar (det, 2)
cadeia determinante = tp.real_para_cadeia(det)
se (txt.posicao_texto(".", determinante, 0) ==-1){
  determinante = determinante + ".00"
}
escreva ("O VALOR DO DETERMINANTE E = ", determinante)
  }
}
