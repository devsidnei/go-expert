package main

import "fmt"

// 10 - Closures
func main() {

	// Função exterior que cria um closure.
	outerFunc := func(x int) func(int) int {
		return func(y int) int {
			return x + y
		}
	}

	// Cria um closure com x igual a 10.
	closure := outerFunc(10)

	// Use o closure para adicionar 5 a 10.
	resultado := closure(5)
	fmt.Println("Resultado:", resultado) // Deve imprimir "Resultado: 15"

}

func div(a int, b int) int {
	return a / b
}
