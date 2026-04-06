programa {
inclua biblioteca Tipos --> tp
inclua biblioteca Texto --> txt
  funcao inicio() {
    
real nota
escreva("")
leia (nota)

cadeia nota_texto = (tp.real_para_cadeia(nota))
se (txt.posicao_texto(".", nota_texto, 0) == -1){
nota_texto = nota_texto + ".0"
}

se (10 >= nota >= 9){
escreva ("NOTA = ", nota_texto, " CONCEITO = A\n")  
}

se (9 > nota >= 7.5){
escreva ("NOTA = ", nota_texto, " CONCEITO = B\n")  
}

se (7.5 > nota >= 6){
escreva ("NOTA = ", nota_texto, " CONCEITO = C\n")  
}

se (0 < nota < 6){
escreva ("NOTA = ", nota_texto, " CONCEITO = D\n")  
}
  }
}
