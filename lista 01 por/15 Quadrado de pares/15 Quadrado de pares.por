programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto --> txt
  inclua biblioteca Tipos--> tp
  funcao inicio() {

cadeia resposta= ""

inteiro valor 
escreva ("")
leia (valor)
inteiro n_pares = valor / 2
para (inteiro i = 1; i <= n_pares; i++) {

inteiro quadrado = (2 * i) * (2 * i)

inteiro par = (2 * i)
cadeia par_texto = tp.inteiro_para_cadeia(par, 10)

cadeia quadrado_texto = tp.inteiro_para_cadeia(quadrado, 10)

resposta =  resposta + par_texto + "^" + par_texto + " = " + quadrado_texto + "\n"
}
escreva (resposta)
  }
}
