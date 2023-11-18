package main

import (
	"fmt"
	"go-expert/matematica/matematica"
)

func main() {

	soma := matematica.Somar(3, 4)
	subtrair := matematica.Subtrair(3, 4)
	dividir := matematica.Dividir(3, 4)
	multiplicar := matematica.Multiplicar(3, 4)

	fmt.Printf("A soma de 3 e 4 é %d\r\n", soma)
	fmt.Printf("A subtração de 3 e 4 é %d\r\n", subtrair)
	fmt.Printf("A divisão de 3 e 4 é %d\r\n", dividir)
	fmt.Printf("A multiplicação de 3 e 4 é %d\r\n", multiplicar)

}
