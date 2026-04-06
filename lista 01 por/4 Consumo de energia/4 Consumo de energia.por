programa {
inclua biblioteca Matematica --> mat
  funcao inicio() {

real salario  
escreva ("")
leia (salario)

real quantidade_kw
escreva ("")
leia (quantidade_kw)

real valor_kw = salario * 0.007
real custo_kw = mat.arredondar (valor_kw, 2)

real valor_pago = valor_kw * quantidade_kw
real custo_consumo = mat.arredondar (valor_pago, 2)

real desconto = valor_pago * 0.9
real custo_desconto = mat.arredondar (desconto, 2)

escreva ("Custo por kW: R$ ", custo_kw, "\n")
escreva ("Custo do consumo: R$ ", custo_consumo, "\n")
escreva ("Custo com desconto: R$ ", custo_desconto , "\n")
  }
}
