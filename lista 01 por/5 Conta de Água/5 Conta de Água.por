programa {
  funcao inicio() {
inteiro conta 
escreva ("")
leia (conta)

real consumo 
escreva ("")
leia (consumo)

cadeia consumidor
escreva ("")
leia (consumidor)

se (consumidor == "r" ou consumidor == "R" ){
real conta_r = (consumo * 0.05) + 5
escreva ("CONTA = ", conta, "VALOR DA CONTA = R$ ", conta_r)
}

se (consumidor == "c" ou consumidor == "C"){
real conta_c = ((consumo - 80) * 0.25 ) + 500
escreva ("CONTA = ", conta, "VALOR DA CONTA = R$ ", conta_c)
}

se (consumidor == "i" ou consumidor == "I"){
real conta_i = ((consumo - 100) * 0.04) + 800
escreva ("CONTA = ", conta, "\nVALOR DA CONTA = R$ ", conta_i)
}
  }
}
