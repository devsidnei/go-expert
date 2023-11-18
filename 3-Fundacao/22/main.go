package main

import (
	"fmt"
	"go-expert/veiculo/veiculo"
)

// Modificadores de acesso => Publico começa com Maisuclo, Privado com Minusculo
func main() {

	carro := veiculo.Carro{}
	carro.SetModelo("Fusca")
	carro.SetMarca("Volkswagen")
	carro.SetAno(1968)
	fmt.Print(veiculo.ImprimirDados(&carro))
}
