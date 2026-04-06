programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto--> txt
  inclua biblioteca Tipos --> tp
  funcao inicio() {

real salario
escreva ("")
leia (salario)

se (salario <= 300){
real reajuste_50 = salario * 1.5
reajuste_50 = mat.arredondar (reajuste_50, 2)

cadeia reajuste_50_texto = tp.real_para_cadeia(reajuste_50)
se ( txt.posicao_texto(".", reajuste_50_texto, 0) == -1){
reajuste_50_texto = reajuste_50_texto + ".00"
}
escreva ("SALARIO COM REAJUSTE = ", reajuste_50_texto, "\n") 
}

se (salario > 300){
real reajuste_30 = salario * 1.3
reajuste_30 = mat.arredondar (reajuste_30, 2)

cadeia reajuste_30_texto = tp.real_para_cadeia(reajuste_30)
se (txt.posicao_texto(".", reajuste_30_texto, 0) == -1){
reajuste_30_texto = reajuste_30_texto + ".00"
}
escreva ("SALARIO COM REAJUSTE = ", reajuste_30, "\n")
}
  }
}
