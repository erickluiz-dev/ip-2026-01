programa {
inclua biblioteca Matematica --> mat
  funcao inicio() {

inteiro casos
escreva ("")
leia (casos)
para (inteiro i = 1; i <= casos; i++) 
{
inteiro publico 
real popular
real geral
real arquibancada
real cadeiras

escreva ("")
leia (publico)

escreva ("")
leia (popular)

escreva ("")
leia (geral)

escreva ("")
leia (arquibancada)

escreva("")
leia (cadeiras)

real renda_popular = (publico * popular) / 100
real renda_geral = (publico * geral / 100) * 5
real renda_arquibancada = (publico * arquibancada / 100) * 10
real renda_cadeiras = (publico * cadeiras / 100) * 20

real renda = renda_popular + renda_geral + renda_arquibancada + renda_cadeiras
mat.arredondar(renda, 2)
escreva ("A RENDA DO JOGO", casos, "E =", renda, "\n" )
}
  }
}
