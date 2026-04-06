 programa {
inclua biblioteca Matematica --> mat
  funcao inicio() {
real n1
escreva ("")
leia (n1)

real n2
escreva ("")
leia (n2)

real n3
escreva ("")
leia (n3)

real soma = (n1) + (n2) + (n3)

real media = soma / 3
real media_final = mat.arredondar(media, 2)

se (media >= 6) {
escreva ("MEDIA = ", media_final, "\nAPROVADO\n")
} senao {
  escreva ("MEDIA = ", media_final, "\nREPROVADO\n")
}
  }
}
