programa {
  inclua biblioteca Matematica --> mat
  inclua biblioteca Texto--> txt
  inclua biblioteca Tipos --> tp  
  funcao inicio() {

real n
escreva ("")
leia (n)
real soma = 0
inteiro dividendo = 0

para (inteiro i = 0; i < n; i++){

dividendo = dividendo + 1
soma = soma + 1 / dividendo 
}
real resultado = mat.arredondar(soma, 6)
escreva (resultado)
  }
}
