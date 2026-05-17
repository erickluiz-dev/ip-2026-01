programa {
inclua biblioteca Matematica --> mat
  funcao inicio() {
real raio
escreva ("")
leia (raio)

real altura
escreva ("")
leia (altura)

real pi = 3.14159 // valor de π

real area_circunferencia = pi * (raio * raio)

real area_retangulo = (2 * pi * raio) * (altura)

real area = 2 * (area_circunferencia) + area_retangulo

real custo = area * 100
real custo_final =mat.arredondar(custo, 2)

escreva ("O VALOR DO CUSTO E = R$ ", custo_final)
  }
}
