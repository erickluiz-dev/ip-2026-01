programa {
inclua biblioteca Tipos --> t
inclua biblioteca Texto --> txt
  funcao inicio() {
cadeia n1
escreva ("")
leia (n1)
se (txt.numero_caracteres(n1) > 1 ou txt.numero_caracteres(n1) <1 ){
escreva ("DIGITO INVALIDO")
retorne }

cadeia n2
escreva ("")
leia (n2)
se (txt.numero_caracteres(n2) > 1 ou txt.numero_caracteres(n2) ){
escreva ("DIGITO INVALIDO")
retorne }

cadeia n3
escreva ("")
leia (n3)
se (txt.numero_caracteres(n3) > 1 ou txt.numero_caracteres(n3)) {
escreva ("DIGITO INVALIDO")
retorne
}
cadeia numero = n1 + n2 + n3

inteiro composicao_inteira = t.cadeia_para_inteiro (numero, 10)

inteiro numero_quadrado = composicao_inteira * composicao_inteira

escreva (composicao_inteira, " ", numero_quadrado)
}
  }
}
