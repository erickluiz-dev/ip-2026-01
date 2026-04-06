programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto --> txt
  inclua biblioteca Tipos--> tp
  funcao inicio() {

real h
escreva ("")
leia (h)

real l
escreva ("")
leia (l)

real raiz_3 = mat.raiz(3, 2.0)

real area= (3 * (l * l) * raiz_3) / 2

real volume = (1 / 3) * h * area
real volume_final = mat.arredondar (volume, 2) 

cadeia resultado = tp.real_para_cadeia (volume_final)
se (txt.posicao_texto(".", resultado, 0) == -1){
resultado = resultado + ".00"
}

escreva ("O VOLUME DA PIRAMIDE E = ", resultado, " METROS CUBICOS\n")

  }
}
