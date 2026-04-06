programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto --> txt
  inclua biblioteca Tipos--> tp
  funcao inicio() {

real numero 
escreva ("")
leia (numero)

real teste_3 = numero / 3
cadeia teste_d3 = tp.real_para_cadeia(teste_3)
se (txt.posicao_texto(".", teste_d3, 0) ==-1){
  escreva ("O NUMERO E DIVISIVEL\n")
retorne
}
cadeia numero_texto = tp.real_para_cadeia (numero)

inteiro tamanho = txt.numero_caracteres(numero_texto)
inteiro ultima_posicao = tamanho - 1

se (txt.obter_caracter(numero_texto, ultima_posicao) ==5 ou txt.obter_caracter(numero_texto, ultima_posicao) ==0 ){
escreva ("O NUMERO E DIVISIVEL\n")
}
senao {
  escreva ("O NUMERO NAO E DIVISIVEL\n")
}




  }
}
