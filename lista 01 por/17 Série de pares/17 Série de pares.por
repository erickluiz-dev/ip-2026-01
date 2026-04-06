programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto--> txt
  inclua biblioteca Tipos --> tp  
  funcao inicio() {

real x
escreva ("")
leia (x)

cadeia resultado = tp.real_para_cadeia(x)

real y
escreva ("")
leia (y)

se ( x % 2 == 0) {
para (inteiro i = 0; i < y; i++){

real numero_ = x + (2 * i)
cadeia numero_par = tp.real_para_cadeia(numero_)

escreva (numero_, " ")
}
}
senao {
escreva ("O PRIMEIRO NUMERO NAO E PAR")
}
escreva ("\n")
  }
}
