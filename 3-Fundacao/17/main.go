package main

import "fmt"

// Ponteiros e Struct
type Conta struct {
	saldo int
}

func newConta() *Conta {
	return &Conta{saldo: 1000}
}

func (c *Conta) usarCartaoDebito(valor int) int {
	c.saldo -= valor
	return c.saldo
}

func (c Conta) usarCartaoCredito(valor int) int {
	c.saldo -= valor
	return c.saldo
}

func main() {
	conta := newConta()

	valorCompra := 10
	conta.usarCartaoDebito(valorCompra)

	fmt.Printf("Compra no Débito no valor de %d - Seu saldo é: %d \n", valorCompra, conta.saldo)

	valorCompra = 134
	conta.usarCartaoCredito(valorCompra)

	fmt.Printf("Compra no Crédito valor de %d - Seu saldo é: %d \n", valorCompra, conta.saldo)

	valorCompra = 14
	conta.usarCartaoDebito(valorCompra)

	fmt.Printf("Compra no Débito valor de %d - Seu saldo é: %d \n", valorCompra, conta.saldo)
}
