programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Tipos --> tp
  inclua biblioteca Texto --> txt
  funcao inicio() {

inteiro repeticoes 
escreva ("")
leia (repeticoes)

cadeia todas_temperaturas = ""

para(inteiro i =1; i <= repeticoes; i++){

real temperatura 
escreva ("")
leia (temperatura)
cadeia temperatura_inicial = tp.real_para_cadeia(temperatura)
se (txt.posicao_texto(".", temperatura_inicial, 0) == -1)
{
temperatura_inicial = temperatura_inicial + ".00"
}

real temperatura_c = ((temperatura - 32) * 5 ) /9.0

real temperatura_final = mat.arredondar(temperatura_c, 2)
cadeia temperatura_celsius = tp.real_para_cadeia(temperatura_final)
se (txt.posicao_texto(".", temperatura_celsius, 0) == -1){
temperatura_celsius = temperatura_celsius + ".00"
}

todas_temperaturas =todas_temperaturas +temperatura_inicial +"\tEQUIVALE A\t" +temperatura_celsius + "\tCELSIUS\n" 
}

escreva (todas_temperaturas)

}
}
