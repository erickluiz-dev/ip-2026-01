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

real delta = (b * b) - ((4* a) *c)
delta = mat.arredondar (delta, 2)
cadeia delta_f = tp.real_para_cadeia(delta)
se (txt.posicao_texto (".", delta_f, 0 ) ==-1){
delta_f = delta_f + ".00"  
}

escreva("O VALOR DE DELTA E = ", delta_f)
  }
}
